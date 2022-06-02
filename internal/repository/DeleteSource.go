package repository

import (
	"context"
	"homework-2/internal/app"
	"homework-2/internal/models"
)

func (r *repository) DeleteSource(ctx context.Context, data models.DCData) (err error) {
	const query1 = `
		select "idSource" from "public.dict"
		where "idUser" = $1 and "idSource" = $2;
	`
	const query2 = `
		delete from "public.dict"
		where "idUser" = $1 and "idSource" = $2;
	`
	resp, err := r.pool.Query(ctx, query1, data.ChatID, data.Source)
	if !resp.Next() {
		err = app.NotFoundErr
		return
	}
	_, err = r.pool.Exec(ctx, query2, data.ChatID, data.Source)

	return
}
