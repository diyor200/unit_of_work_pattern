package repository

import (
	"context"
	"github.com/diyor200/uof/internal/uow"
	"log"

	"github.com/diyor200/uof/internal/domains"
)

type UserRepo interface {
	AddUser(ctx context.Context, data domains.User) (domains.User, error)
	ChangeStatus(ctx context.Context, data domains.User) error
}

type userRepo struct {
	tx uow.Tx
}

func NewRepos(tx uow.Tx) UserRepo {
	return &userRepo{tx: tx}
}

func (r *userRepo) AddUser(ctx context.Context, data domains.User) (domains.User, error) {
	err := r.tx.QueryRow(ctx, "insert into users(name, email, status) values ($1, $2, $3) returning id",
		data.Name, data.Email, data.Status).Scan(&data.ID)
	if err != nil {
		log.Println(err)
		return domains.User{}, err
	}

	return data, nil
}

func (r *userRepo) ChangeStatus(ctx context.Context, data domains.User) error {
	_,
		err := r.tx.Exec(ctx, "update users set status = $1 where id = $2", data.Status, data.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
