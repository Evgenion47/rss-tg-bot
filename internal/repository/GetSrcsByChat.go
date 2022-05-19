package repository

import (
	"context"
	"homework-2/internal/app"
	"homework-2/internal/models"
)

func (r *repository) GetSrcsByChat(ctx context.Context, user models.User) (srcs models.Sources, err error) {

	const query = `
		select "idSource" from "public.dict" where "idUser" = $1;
	`
	rows, err := r.pool.Query(ctx, query, user.ChatID)
	if err != nil {
		err = app.NotFoundErr
		return
	}
	defer rows.Close()

	for rows.Next() {
		var foo string
		if err = rows.Scan(&foo); err != nil {
			return
		}
		srcs.Sources = append(srcs.Sources, foo)
	}

	return
}
