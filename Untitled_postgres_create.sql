CREATE TABLE "public.Users" (
	"id" serial NOT NULL,
	"telegramNickname" VARCHAR(255) NOT NULL,
	CONSTRAINT "Users_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.Sources" (
	"id_source" serial NOT NULL,
	"link" VARCHAR(255) NOT NULL,
	CONSTRAINT "Sources_pk" PRIMARY KEY ("id_source")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.dict" (
	"id" serial NOT NULL,
	"id_source" integer NOT NULL,
	"id_user" integer NOT NULL,
	CONSTRAINT "dict_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);





ALTER TABLE "dict" ADD CONSTRAINT "dict_fk0" FOREIGN KEY ("id_source") REFERENCES "Sources"("id_source");
ALTER TABLE "dict" ADD CONSTRAINT "dict_fk1" FOREIGN KEY ("id_user") REFERENCES "Users"("id");




