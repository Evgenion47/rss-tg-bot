package repository

import (
	"context"
	"homework-2/internal/models"
)

func (r *repository) CreateUser(ctx context.Context, user models.User) (err error) {

	const query = `
		insert into "public.Users" ("idUser") 
		VALUES ($1);
	`

	_, err = r.pool.Exec(ctx, query, user.ChatID)

	return

}
