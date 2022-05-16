-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table "public.Users"
(
    "idUser" integer NOT NULL,
    CONSTRAINT "Users_pk" PRIMARY KEY ("idUser")
);

create table "public.Source"
(
    "idSource"   VARCHAR(255) NOT NULL,
    CONSTRAINT "Source_pk" PRIMARY KEY ("idSource")
);

create table "public.dict"
(
    "idUser"   integer      NOT NULL,
    "idSource" VARCHAR(255) NOT NULL,
    "lastUpdate" TIMESTAMP  NOT NULL,
    unique ("idUser", "idSource", "lastUpdate")
);

alter table "public.dict"
    add constraint "dict_fk0" foreign key ("idUser") references "public.Users" ("idUser");
alter table "public.dict"
    add constraint "dict_fk1" foreign key ("idSource") references "public.Source" ("idSource");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists "public.dict";
drop table if exists "public.Source";
drop table if exists "public.Users";
-- +goose StatementEnd
