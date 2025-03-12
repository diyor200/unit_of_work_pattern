package repository

import (
	"context"
	"log"

	"github.com/diyor200/uof/internal/domains"
	"github.com/jackc/pgx/v4"
)

type UserRepo interface {
	AddUser(ctx context.Context, data domains.User) (domains.User, error)
	ChangeStatus(ctx context.Context, data domains.User) error
}

type Repos struct {
	db *pgx.Conn
}

func NewRepos(db *pgx.Conn) *Repos {
	return &Repos{db: db}
}

func (r *Repos) AddUser(ctx context.Context, data domains.User) (domains.User, error) {
	err := r.db.QueryRow(ctx, "insert into users(name, email, status) values ($1, $2, $3) returning id",
		data.Name, data.Email, data.Status).Scan(&data.ID)
	if err != nil {
		log.Println(err)
		return domains.User{}, err
	}

	return data, nil
}

func (r *Repos) ChangeStatus(ctx context.Context, data domains.User) error {
	_,
		err := r.db.Exec(ctx, "update users set status = $1 where id = $2", data.Status, data.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
