package uow

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UOWManager interface {
	New(ctx context.Context) (UOW, error)
}

type Tx interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type UOW interface {
	Commit() error
	Rollback() error
	GetTx() Tx
}

type uowManager struct {
	pool *pgxpool.Pool
}

type uow struct {
	tx pgx.Tx
}

func NewUOWManager(pool *pgxpool.Pool) UOWManager {
	return &uowManager{pool: pool}
}

func (m *uowManager) New(ctx context.Context) (UOW, error) {
	tx, err := m.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &uow{tx: tx}, nil
}

func (u *uow) Commit() error {
	return u.tx.Commit(context.Background())
}

func (u *uow) Rollback() error {
	return u.tx.Rollback(context.Background())
}

func (u *uow) GetTx() Tx {
	return u.tx
}
