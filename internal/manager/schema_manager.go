package manager

import (
	"context"
	"fmt"
	"sort"
	"strings"

	schemapilotv1 "github.com/gjduggins/schemapilot-operator/api/v1"
	"github.com/gjduggins/schemapilot-operator/internal/db"
	"github.com/gjduggins/schemapilot-operator/internal/model"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

type SchemaManager struct {
	DA db.DatabaseAdapter
	SC db.ScmAdapter
}

func NewSchemaManager(da db.DatabaseAdapter, sc db.ScmAdapter) *SchemaManager {
	return &SchemaManager{DA: da, SC: sc}
}

func (sm *SchemaManager) ApplySchemaChanges(ctx context.Context, schemapilot *schemapilotv1.SchemaPilot) error {
	log := logf.FromContext(ctx)
	if schemapilot == nil || len(schemapilot.Spec.SchemaFiles) == 0 {
		log.Info("CR not fully populated yet, skipping reconcile")
		return nil
	}

	desiredOrder := []string{"Table", "Index", "PackageHeader", "View", "Procedure", "Function", "Trigger", "Grant", "Synonym", "PackageBody", "Other"}

	// Create a map for quick lookup of priority
	orderMap := make(map[string]int)
	for i, obj := range desiredOrder {
		orderMap[obj] = i
	}

	// Sort the slice in-place
	sort.Slice(schemapilot.Spec.SchemaFiles, func(i, j int) bool {
		typeI := schemapilot.Spec.SchemaFiles[i].ObjectType
		typeJ := schemapilot.Spec.SchemaFiles[j].ObjectType
		return orderMap[typeI] < orderMap[typeJ]
	})

	for _, schemaFile := range schemapilot.Spec.SchemaFiles {
		if schemaFile.Name == "" || schemaFile.Version == "" {
			log.Info("Schema file entry incomplete, skipping", "file", schemaFile.Name, "version", schemaFile.Version)
			continue
		}
		log.Info("***** STARTING " + schemaFile.Name + " *****")

		apply, err := sm.applyChange(ctx, schemaFile.Name, schemaFile.Version)
		if err != nil {
			log.Error(err, "error checking applied version", "file", schemaFile.Name)
			return err
		}

		if !apply {
			log.Info("Schema file already applied, skipping", "file", schemaFile.Name, "version", schemaFile.Version)
			continue
		}

		log.Info("Fetching schema file from Github", "file", schemaFile.Name, "version", schemaFile.Version)
		script, err := sm.SC.FetchFile(schemaFile.Name, schemaFile.Version)

		if err != nil {
			log.Error(err, "failed to read schema file", "file", schemaFile.Name)
			return err
		}

		sqlCommand := string(script)
		log.Info("sqlCommand for " + schemaFile.Name + ": " + sqlCommand)
		if schemaFile.ObjectType == "Table" {
			sqlCommand, err = sm.CreateTableCommand(ctx, sqlCommand)
			if err != nil {
				log.Error(err, "failed to create table command", "file", schemaFile.Name)
				return err
			}
			if sqlCommand == "" {
				log.Info("No changes to apply for table, skipping", "file", schemaFile.Name)
				if rows, err := sm.DA.UpdateSchemaVersion(ctx, schemaFile.Name, schemaFile.Version); err != nil {
					log.Error(err, "failed to update schema version", "file", schemaFile.Name, "version", schemaFile.Version)
					return err
				} else {
					log.Info("Updated schema version", "file", schemaFile.Name, "version", schemaFile.Version, "rowsAffected", rows)
				}
				continue
			}
		}

		log.Info("Applying schema file", "file", schemaFile.Name, "version", schemaFile.Version)
		if err := sm.DA.ApplyScript(ctx, sqlCommand); err != nil {
			log.Error(err, "failed to apply schema file", "file", schemaFile.Name)
			return err
		}

		if rows, err := sm.DA.UpdateSchemaVersion(ctx, schemaFile.Name, schemaFile.Version); err != nil {
			log.Error(err, "failed to update schema version", "file", schemaFile.Name, "version", schemaFile.Version)
			return err
		} else {
			log.Info("Updated schema version", "file", schemaFile.Name, "version", schemaFile.Version, "rowsAffected", rows)
		}
	}

	return nil
}

func (sm *SchemaManager) applyChange(ctx context.Context, script string, version string) (bool, error) {
	appliedVersion, err := sm.DA.GetAppliedVersion(ctx, script)

	if err != nil {
		return false, fmt.Errorf("failed to get applied version for script %s: %w", script, err)
	}

	if appliedVersion == version {
		return false, nil
	}

	return true, nil
}

func (sm *SchemaManager) CreateTableCommand(ctx context.Context, sqlCommand string) (string, error) {
	log := logf.FromContext(ctx)
	tbl, err := ParseSQL(sqlCommand)
	if err != nil {
		log.Error(err, "failed to parse schema file SQL")
		return "", err
	}

	log.Info("Parsed table", "name", tbl.Name, "columns", tbl.Columns)
	existingTable, err := sm.DA.CheckForExistingTable(ctx, tbl.Name)
	if err != nil {
		log.Error(err, "failed to check for existing table", "table", tbl.Name)
		return "", err
	}

	if existingTable != nil {
		log.Info("Table already exists in database, Checking Modifications", "table", tbl.Name)
		addColumns, modifiedColumns := CompareColumns(tbl.Columns, existingTable.Columns)
		if len(addColumns) == 0 && len(modifiedColumns) == 0 {
			log.Info("No changes detected for table, skipping ALTER", "table", tbl.Name)
			return "", nil
		}
		sqlCommand, err = sm.DA.GenerateAlterScript(tbl.Name, addColumns, modifiedColumns)
		if err != nil {
			log.Error(err, "failed to generate ALTER script", "table", tbl.Name)
			return "", err
		}
		log.Info("Alters script " + sqlCommand)
	}
	return sqlCommand, nil
}

func CompareColumns(a, b []model.Column) (addColumns, modifiedColumns []model.Column) {
	// Build a lookup map from b
	bMap := make(map[string]string)
	for _, col := range b {
		bMap[strings.ToUpper(col.Name)] = col.Type
	}

	// Compare each column in a with b
	for _, col := range a {
		bType, exists := bMap[strings.ToUpper(col.Name)]
		if !exists {
			addColumns = append(addColumns, col)
			continue
		}
		if !strings.EqualFold(bType, col.Type) {
			modifiedColumns = append(modifiedColumns, col)
		}
	}

	return addColumns, modifiedColumns
}
