package repository

import (
	"context"
	"homework-2/internal/app"
	"homework-2/internal/models"
)

func (r *repository) CreateUser(ctx context.Context, user models.User) (err error) {
	const query1 = `
		select "idUser" from "public.Users"
		where "idUser" = $1;
	`
	const query2 = `
		insert into "public.Users" ("idUser") 
		VALUES ($1);
`
	resp, err := r.pool.Query(ctx, query1, user.ChatID)
	if resp.Next() {
		err = app.AlreadyExistErr
		return
	}
	_, err = r.pool.Exec(ctx, query2, user.ChatID)
	return
}
