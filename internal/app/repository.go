package app

import (
	"context"
	"errors"
	"homework-2/internal/models"
)

var NotFoundErr = errors.New("Bad luck, it's already exist")
var AlreadyExistErr = errors.New("Bad luck, it's already exist")

type Repository interface {
	CreateUser(ctx context.Context, user models.User) error
	CreateSource(ctx context.Context, data models.DCData) error
	DeleteSource(ctx context.Context, data models.DCData) error
	GetSrcsByChat(ctx context.Context, user models.User) (models.Sources, error)
	GetRSSBySource(ctx context.Context, data models.DCData) (models.RSSNews, error)
}
