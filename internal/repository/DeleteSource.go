package repository

import (
	"context"
	"homework-2/internal/models"
)

func (r *repository) DeleteSource(ctx context.Context, data models.DCData) (err error) {

	const query = `
		delete from "public.dict"
		where "idUser" = $1 and "idSource" = $2;
	`

	_, err = r.pool.Exec(ctx, query, data.ChatID, data.Source)

	return

}
