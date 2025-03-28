package arreglo

import (
	"context"
	"recuCorte1/src/user/domain/model"
)

func (r *UserRepository) GetAll(ctx context.Context) ([]*model.User, error) {

	var users []*model.User

	for i := range r.users {
		users = append(users, &r.users[i])
	}

	return users, nil
}
