package arreglo

import (
	"context"
	"errors"
	"recuCorte1/src/user/domain/model"
)

func (r *UserRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}
