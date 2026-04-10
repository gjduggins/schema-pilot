package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/gjduggins/schemapilot-operator/internal/model"
	_ "github.com/sijms/go-ora/v2"
)

type OracleAdapter struct {
	DB *sql.DB
}

// Connect creates a new Oracle database connection
func NewOracleAdapter(ctx context.Context, dsn string) (*OracleAdapter, error) {

	db, err := sql.Open("oracle", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open Oracle connection: %w", err)
	}

	// Test the connection
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping Oracle database: %w", err)
	}

	return &OracleAdapter{DB: db}, nil
}

// Close closes the database connection
func (oa *OracleAdapter) Close() error {
	return oa.DB.Close()
}

func (oa *OracleAdapter) GetAppliedVersion(ctx context.Context, script string) (string, error) {
	var version string

	query := "select script_version from schema_pilot_history where script_name = :1 order by applied_at desc fetch first 1 rows only"

	err := oa.DB.QueryRowContext(ctx, query, script).Scan(&version)
	if err != nil {
		if err == sql.ErrNoRows {
			// Not found is valid — return empty string or custom handling
			return "", nil
		}
		return "", fmt.Errorf("query failed: %w", err)
	}

	return version, nil
}

// ExecuteUpdate executes an INSERT, UPDATE, or DELETE query
func (oa *OracleAdapter) UpdateSchemaVersion(ctx context.Context, name string, version string) (int64, error) {

	query := "INSERT INTO schema_pilot_history (script_name, script_version, applied_at) VALUES (:1, :2, SYSTIMESTAMP)"
	result, err := oa.DB.ExecContext(ctx, query, name, version)

	if err != nil {
		return 0, fmt.Errorf("failed to execute update: %w", err)
	}
	return result.RowsAffected()
}

// ApplyDDLScript reads a DDL script from a file and executes it
func (oa *OracleAdapter) ApplyScript(ctx context.Context, sqlCommand string) error {
	// Prepend /scripts/ to the file path
	// statements := splitStatements(sqlCommand)

	// for _, stmt := range statements {
	// 	if stmt == "" {
	// 		continue
	// 	}
	// 	_, err := oa.DB.ExecContext(ctx, stmt)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to execute DDL statement: %w\nStatement: %s", err, stmt)
	// 	}
	// }

	_, err := oa.DB.ExecContext(ctx, sqlCommand)
	if err != nil {
		return fmt.Errorf("failed to execute DDL statement: %w\nStatement: %s", err, sqlCommand)
	}
	return nil
}

func (oa *OracleAdapter) CheckForExistingTable(ctx context.Context, tableName string) (*model.Table, error) {
	rows, err := oa.DB.Query(`
        SELECT column_name, 
		       decode(data_type,'VARCHAR2', data_type || '(' || data_length || ')',
				                'NUMBER', data_type || decode(data_precision,null,null,'(' || data_precision || decode(data_scale,0,null,null,null,',' || data_scale) || ')'),
								data_type) as data_type
        FROM all_tab_cols
        WHERE owner = USER AND table_name = :1
        ORDER BY column_id
    `, tableName)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []model.Column
	for rows.Next() {
		var col model.Column
		if err := rows.Scan(&col.Name, &col.Type); err != nil {
			return nil, err
		}
		columns = append(columns, col)
	}

	// If no columns were found, the table doesn't exist
	if len(columns) == 0 {
		return nil, nil
	}

	return &model.Table{
		Name:    tableName,
		Columns: columns,
	}, nil
}

func (oa *OracleAdapter) GenerateAlterScript(tableName string, addColumns, modifyColumns []model.Column) (string, error) {
	var alters []string

	var colDefs []string
	for _, col := range addColumns {
		if col.Name == "" || col.Type == "" {
			continue
		}
		colDefs = append(colDefs, fmt.Sprintf("%s %s", col.Name, col.Type))
	}

	var colMods []string
	for _, col := range modifyColumns {
		if col.Name == "" || col.Type == "" {
			continue
		}
		colMods = append(colMods, fmt.Sprintf("%s %s", col.Name, col.Type))
	}

	if len(colDefs) > 0 || len(colMods) > 0 {
		alter := fmt.Sprintf("ALTER TABLE %s ", tableName)
		alters = append(alters, alter)
		if len(colDefs) > 0 {
			alter := fmt.Sprintf("ADD ( %s )", strings.Join(colDefs, ", "))
			alters = append(alters, alter)
		}
		if len(colMods) > 0 {
			alter := fmt.Sprintf("MODIFY ( %s )", strings.Join(colMods, ", "))
			alters = append(alters, alter)
		}
		alters = append(alters, ";")
	}

	return strings.Join(alters, "\n"), nil
}

// splitStatements splits a SQL script into individual statements
func splitStatements(script string) []string {
	var statements []string
	var current string

	for _, char := range script {
		current += string(char)
		if char == ';' {
			statements = append(statements, current[:len(current)-1]) // Remove semicolon
			current = ""
		}
	}

	if current != "" {
		statements = append(statements, current)
	}

	return statements
}
