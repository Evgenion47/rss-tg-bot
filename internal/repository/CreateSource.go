package repository

import (
	"context"
	"homework-2/internal/models"
)

func (r *repository) CreateSource(ctx context.Context, data models.DCData) (err error) {

	const query = `
		select createSource($1, $2);
	`
	_, err = r.pool.Exec(ctx, query, data.ChatID, data.Source)

	return
}

//create function createsource(integer, character varying) returns void
//language plpgsql
//as
//$$
//DECLARE
//BEGIN
//IF (select count("idSource") from "public.Source") = 0 THEN
//insert into "public.Source"("idSource", "lastUpdate") VALUES ($2,0);
//END IF;
//insert into "public.dict"("idUser", "idSource") VALUES ($1, $2);
//END
//$$;
//
//alter function createsource(integer, varchar) owner to postgres;
