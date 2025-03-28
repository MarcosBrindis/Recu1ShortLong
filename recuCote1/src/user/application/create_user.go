package application

import (
	"context"
	"recuCorte1/src/user/domain/model"
	"recuCorte1/src/user/domain/ports"
)

type CreateUserUsecase struct {
	Repository ports.UserRepository
}

func (uc *CreateUserUsecase) Execute(ctx context.Context, user *model.User) error {
	return uc.Repository.Create(ctx, user)
}
