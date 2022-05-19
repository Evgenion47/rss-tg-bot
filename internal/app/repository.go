package app

import (
	"context"
	"errors"
	"homework-2/internal/models"
)

var NotFoundErr = errors.New("Bad luck, nothing found")
var AlreadyExistErr = errors.New("Bad luck, it's already exist")
var HavntUpd = errors.New("Bad luck, have not updates")
var WrongLink = errors.New("Bad luck, wrong link")
var UserDoesNotExist = errors.New("Bad luck, user does not exist")

type Repository interface {
	CreateUser(ctx context.Context, user models.User) error
	CreateSource(ctx context.Context, data models.DCData) error
	DeleteSource(ctx context.Context, data models.DCData) error
	GetSrcsByChat(ctx context.Context, user models.User) (models.Sources, error)
	GetRSSBySource(ctx context.Context, data models.DCData) (models.RSSNews, error)
}
