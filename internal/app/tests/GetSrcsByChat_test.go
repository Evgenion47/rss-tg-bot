package tests

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"homework-2/internal/app"
	"homework-2/internal/models"
	"homework-2/pkg/api"
	"testing"
)

func TestGetSrcsByChat(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mrepo := app.NewRepositoryMock(mc)
	mrepo.GetSrcsByChatMock.Return(models.Sources{Sources: []string{}}, nil)
	service := app.New(mrepo)

	resp, err := service.GetSrcsByChat(context.Background(), &api.ChData{ChatID: 1})
	assert.Nil(t, err)
	assert.Equal(t, resp.Sources, []string{})
}
