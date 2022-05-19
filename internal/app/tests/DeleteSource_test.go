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

func TestDeleteSource(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mrepo := app.NewRepositoryMock(mc)
	mrepo.DeleteSourceMock.Return(nil)
	service := app.New(mrepo)

	_, err := service.DeleteSource(context.Background(), &api.ChSrcData{ChatID: 1, Source: "http://blog.golang.org/feed.atom"})
	assert.Nil(t, err)
}

func TestDeleteSourceDublicate(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mrepo := app.NewRepositoryMock(mc)
	mrepo.DeleteSourceMock.Return(app.NotFoundErr)
	service := app.New(mrepo)

	service.DeleteSource(context.Background(), &api.ChSrcData{ChatID: 1, Source: "http://blog.golang.org/feed.atom"})
	_, err := service.DeleteSource(context.Background(), &api.ChSrcData{ChatID: 1, Source: "http://blog.golang.org/feed.atom"})
	assert.Equal(t, err, errors.New("Bad luck, nothing found"))
}
