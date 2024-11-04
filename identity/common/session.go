package common

import (
	"context"
	"database/sql"
)

type Session struct {
	db        *sql.DB
	txOptions *sql.TxOptions
	ctx       context.Context
}

func NewSession(db *sql.DB, txOptions *sql.TxOptions, ctx context.Context) *Session {
	return &Session{
		db:        db,
		txOptions: txOptions,
		ctx:       ctx,
	}
}

func (s *Session) Transaction(ctx context.Context, fn func(context.Context) error) error {
	tx, err := s.db.BeginTx(ctx, s.txOptions)
	if err != nil {
		return err
	}
	err = fn(context.WithValue(ctx, "db", tx))
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func CreateSerializableTxOption() *sql.TxOptions {
	return &sql.TxOptions{Isolation: sql.LevelSerializable}
}

func (s *Session) ExecQuery(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if tx, ok := ctx.Value("db").(*sql.Tx); ok {
		result, err := tx.ExecContext(ctx, query, args...)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	result, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Session) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if tx, ok := ctx.Value("db").(*sql.Tx); ok {
		return tx.QueryRowContext(ctx, query, args...)
	}
	return s.db.QueryRowContext(ctx, query, args...)
}

func (s *Session) QueryMultiRows(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if tx, ok := ctx.Value("db").(*sql.Tx); ok {
		return tx.QueryContext(ctx, query, args...)
	}
	return s.db.QueryContext(ctx, query, args...)
}
