package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Transaction interface {
	Querier
	ExecTxn(ctx context.Context, execuFn func(q *Queries) error) error
}

type transaction struct {
	db *sql.DB
	*Queries
}

func NewTransaction(db *sql.DB) *transaction {
	return &transaction{
		db: db,
		// Queries: New(db),
	}
}

func (t *transaction) ExecTxn(ctx context.Context, execuFn func(q *Queries) error) error {
	sqlTxn, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("couldn't initiate txn: %s", err)
	}
	qr := New(sqlTxn)
	err = execuFn(qr)
	if err != nil {
		if err = sqlTxn.Rollback(); err != nil {
			return fmt.Errorf("couldn't complete txn: %s", err)
		}
	}
	return sqlTxn.Commit()
}
