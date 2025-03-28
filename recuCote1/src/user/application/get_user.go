package application

import (
	"context"
	"recuCorte1/src/user/domain/model"
	"recuCorte1/src/user/domain/ports"
)

type GetUserUsecase struct {
	Repository ports.UserRepository
}

func (uc *GetUserUsecase) Execute(ctx context.Context, id int) (*model.User, error) {
	return uc.Repository.GetByID(ctx, id)
}
