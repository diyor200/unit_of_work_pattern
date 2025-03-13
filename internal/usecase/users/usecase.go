package users

import (
	"context"
	"github.com/diyor200/uof/internal/domains"
	"github.com/diyor200/uof/internal/repository"
	"github.com/diyor200/uof/internal/uow"
)

type Usecase struct {
	uowManager uow.UOWManager // Inject factory to create UOW
}

func New(uowManager uow.UOWManager) *Usecase {
	return &Usecase{uowManager: uowManager}
}

func (u *Usecase) AddUser(ctx context.Context, data domains.User) (domains.User, error) {
	// Start UOW (transaction) for this request
	tx, err := u.uowManager.New(ctx)
	if err != nil {
		return domains.User{}, err
	}

	// Always ensure rollback unless commit succeeds
	defer tx.Rollback()

	// Repositories use transaction
	userRepo := repository.NewRepos(tx.GetTx())

	// Add user
	user, err := userRepo.AddUser(ctx, data)
	if err != nil {
		return domains.User{}, err
	}

	// Update status to active
	user.Status = domains.UserStatusActive
	err = userRepo.ChangeStatus(ctx, user) // you should use user here, not data
	if err != nil {
		return domains.User{}, err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return domains.User{}, err
	}

	return user, nil
}
