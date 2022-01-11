package database

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	config "github.com/isaqueveras/authentication-microservice/configuration"
	_ "github.com/lib/pq"
)

// DBTransaction used to aggregate transactions
type DBTransaction struct {
	postgres *sql.Tx
	Builder  squirrel.StatementBuilderType
	ctx      context.Context
}

// OpenConnection initialize connection with database
func OpenConnection(ctx context.Context, readOnly bool) (*DBTransaction, error) {
	var (
		t   = &DBTransaction{}
		db  *sql.DB
		err error
	)

	if db, err = sql.Open(config.Get().Database.Driver, config.Get().Database.Url); err != nil {
		return t, err
	}

	defer db.Close()

	transaction, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  readOnly,
	})

	if err != nil {
		return nil, err
	}

	t.ctx = ctx
	t.postgres = transaction
	t.Builder = squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		RunWith(t.postgres)

	return t, nil
}

// Commit commit pending transactions for all open databases
func (t *DBTransaction) Commit() (erro error) {
	return t.postgres.Commit()
}

// Rollback rollback transaction pending for all open databases
func (t *DBTransaction) Rollback() {
	_ = t.postgres.Rollback()
}
