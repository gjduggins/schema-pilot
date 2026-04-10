package db

import (
	"context"
	"fmt"
	"strings"
)

func Connect(ctx context.Context, dsn string) (DatabaseAdapter, error) {

	switch {
	case strings.HasPrefix(dsn, "oracle://"):
		return NewOracleAdapter(ctx, dsn)
	default:
		return nil, fmt.Errorf("unknown database in dsn: %s", dsn)
	}
}
