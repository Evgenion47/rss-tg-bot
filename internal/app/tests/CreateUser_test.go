package tests

import (
	"context"
	"errors"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"homework-2/internal/app"
	"homework-2/pkg/api"
	"testing"
)

func TestCreateUser(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mrepo := app.NewRepositoryMock(mc)
	mrepo.CreateUserMock.Return(nil)
	service := app.New(mrepo)

	_, err := service.CreateUser(context.Background(), &api.ChData{ChatID: 1})
	assert.Nil(t, err)
}

func TestCreateUserDublicate(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mrepo := app.NewRepositoryMock(mc)
	mrepo.CreateUserMock.Return(app.AlreadyExistErr)
	service := app.New(mrepo)

	_, err := service.CreateUser(context.Background(), &api.ChData{ChatID: 1})
	_, err = service.CreateUser(context.Background(), &api.ChData{ChatID: 1})
	assert.Equal(t, err, errors.New("Bad luck, it's already exist"))
}
