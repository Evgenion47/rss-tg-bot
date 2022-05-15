-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table "public.Users"
(
    "idChat" serial NOT NULL,
    CONSTRAINT "Users_pk" PRIMARY KEY ("idChat")
);

create table "public.Source"
(
    "idSource"   VARCHAR(255) NOT NULL,
    "lastUpdate" TIMESTAMP    NOT NULL,
    CONSTRAINT "Source_pk" PRIMARY KEY ("idSource")
);

create table "public.dict"
(
    "idUser"   integer      NOT NULL,
    "idSource" VARCHAR(255) NOT NULL
);

alter table "public.dict"
    add constraint "dict_fk0" foreign key ("idUser") references "public.Users" ("idChat");
alter table "public.dict"
    add constraint "dict_fk1" foreign key ("idSource") references "public.Source" ("idSource");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table "public.dict"
drop table "public.Source"
drop table "public.Users"
-- +goose StatementEnd
