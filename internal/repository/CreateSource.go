package repository

import (
	"context"
	"homework-2/internal/app"
	"homework-2/internal/models"
)

func (r *repository) CreateSource(ctx context.Context, data models.DCData) (err error) {
	const query1 = `
		select "idUser" from "public.Users" 
		where "idUser" = $1;
`
	const query2 = `
		select "source" from "public.dict"
		where "idUser" = $1 and "idSource" = $2;
`
	resp, err := r.pool.Query(ctx, query1, data.ChatID)
	if !resp.Next() {
		err = app.UserDoesNotExist
		return
	}
	resp, err = r.pool.Query(ctx, query2, data.ChatID, data.Source)
	if resp.Next() {
		err = app.AlreadyExistErr
		return
	}

	const query3 = `
		select createSource($1, $2);
	`
	_, err = r.pool.Exec(ctx, query3, data.ChatID, data.Source)

	return
}
