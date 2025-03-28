package ports

import (
	"context"
	"recuCorte1/src/user/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
}
