package users

import (
	"context"

	"github.com/diyor200/uof/internal/domains"
	"github.com/diyor200/uof/internal/repository"
)

type UserUseCase interface {
	AddUser(ctx context.Context, data domains.User) (domains.User, error)
}

type Usecase struct {
	repo repository.UserRepo
}

func New(repo repository.UserRepo) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) AddUser(ctx context.Context, data domains.User) (domains.User, error) {
	// add user
	user, err := u.repo.AddUser(ctx, data)
	if err != nil {
		return domains.User{}, err
	}

	// update status
	err = u.repo.ChangeStatus(ctx, data)
	if err != nil {
		return domains.User{}, err
	}

	return user, nil
}
