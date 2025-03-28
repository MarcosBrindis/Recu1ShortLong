package application

import (
	"context"
	"recuCorte1/src/user/domain/model"
	"recuCorte1/src/user/domain/ports"
)

type GetAllUsersUsecase struct {
	Repository ports.UserRepository
}

func (uc *GetAllUsersUsecase) Execute(ctx context.Context) ([]*model.User, error) {
	return uc.Repository.GetAll(ctx)
}
