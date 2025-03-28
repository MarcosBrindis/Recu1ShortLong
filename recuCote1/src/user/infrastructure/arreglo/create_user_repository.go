package arreglo

import (
	"context"
	"recuCorte1/src/user/domain/model"
)

type UserRepository struct {
	users []model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []model.User{},
	}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	user.ID = len(r.users) + 1

	r.users = append(r.users, *user)
	return nil
}
