package tests

import (
	"context"
	"errors"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"homework-2/internal/app"
	"homework-2/internal/models"
	"homework-2/pkg/api"
	"testing"
)

func TestGetRSSBySource(t *testing.T) {

	res := []string{
		"2022-04-19T00:00:00+00:00 : Go Developer Survey 2021 Results",
		"2022-04-12T00:00:00+00:00 : When To Use Generics",
		"2022-04-05T00:00:00+00:00 : Get familiar with workspaces",
		"2022-03-31T00:00:00+00:00 : How Go Mitigates Supply Chain Attacks",
		"2022-03-22T00:00:00+00:00 : An Introduction To Generics",
		"2022-03-15T00:00:00+00:00 : Go 1.18 is released!",
		"2022-01-31T00:00:00+00:00 : Announcing Go 1.18 Beta 2",
		"2022-01-14T00:00:00+00:00 : Two New Tutorials for 1.18",
		"2021-12-14T00:00:00+00:00 : Go 1.18 Beta 1 is available, with generics",
		"2021-11-10T00:00:00+00:00 : Twelve Years of Go",
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	mrepo := app.NewRepositoryMock(mc)
	mrepo.GetRSSBySourceMock.Return(models.RSSNews{News: res}, nil)
	service := app.New(mrepo)

	resp, err := service.GetRSSBySource(context.Background(), &api.ChSrcData{ChatID: 1, Source: "http://blog.golang.org/feed.atom"})
	assert.Equal(t, resp.News, res)
	assert.Nil(t, err)
}

func TestGetRSSBySourceHavntUpd(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mrepo := app.NewRepositoryMock(mc)
	mrepo.GetRSSBySourceMock.Return(models.RSSNews{[]string{}}, app.HavntUpd)
	service := app.New(mrepo)

	service.GetRSSBySource(context.Background(), &api.ChSrcData{ChatID: 1, Source: "http://blog.golang.org/feed.atom"})
	resp, err := service.GetRSSBySource(context.Background(), &api.ChSrcData{ChatID: 1, Source: "http://blog.golang.org/feed.atom"})
	assert.Equal(t, err, errors.New("Bad luck, have not updates"))
	assert.Equal(t, resp.News, []string{})
}

func TestGetRSSBySourceWrongLink(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mrepo := app.NewRepositoryMock(mc)
	mrepo.GetRSSBySourceMock.Return(models.RSSNews{[]string{}}, app.WrongLink)
	service := app.New(mrepo)

	resp, err := service.GetRSSBySource(context.Background(), &api.ChSrcData{ChatID: 1, Source: "og.golang.org/feed.atom"})

	assert.Equal(t, resp.News, []string{})
	assert.Equal(t, err, errors.New("Bad luck, wrong link"))
}
