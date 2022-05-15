package repository

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"homework-2/internal/models"
)

func (r *repository) CreateProduct(ctx context.Context, user models.User) (empt *empty.Empty, err error) {

	const query = `
		insert into Users (
			idChat
		) VALUES (
			$1
		)
	`

	cmd, err := r.pool.Exec(ctx, query, user.UserID)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return

}
