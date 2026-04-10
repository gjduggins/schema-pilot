package db

import (
	"context"

	"github.com/gjduggins/schemapilot-operator/internal/model"
)

type DatabaseAdapter interface {
	GetAppliedVersion(ctx context.Context, script string) (string, error)
	ApplyScript(ctx context.Context, sql string) error
	UpdateSchemaVersion(ctx context.Context, script string, version string) (int64, error)
	CheckForExistingTable(ctx context.Context, tableName string) (*model.Table, error)
	GenerateAlterScript(tableName string, addColumns, modifyColumns []model.Column) (string, error)
	Close() error
}
