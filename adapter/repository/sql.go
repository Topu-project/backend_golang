package repository

import "context"

type SQL interface {
	ExecuteContext(string, ...any) error
	Query(string, ...any) (Rows, error)
	QueryContext(context.Context, string, ...any) (Rows, error)
	QueryRow(string, ...any) Row
	QueryRowContext(context.Context, string, ...any) Row
	BeginTx(ctx context.Context) (Tx, error)
}

type Rows interface {
	Scan(dest ...any) error
	Next() bool
	Err() error
	Close() error
}

type Row interface {
	Scan(dest ...any) error
}

type Tx interface {
	ExecuteContext(context.Context, string, ...any) error
	QueryContext(context.Context, string, ...any) (Rows, error)
	QueryRowContext(context.Context, string, ...any) Row
	Commit() error
	Rollback() error
}
