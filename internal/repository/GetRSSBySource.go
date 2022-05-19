package repository

import (
	"context"
	"fmt"
	"homework-2/internal/app"
	"homework-2/internal/models"
	"homework-2/rss"
	"path/filepath"
	"time"
)

func (r *repository) GetRSSBySource(ctx context.Context, data models.DCData) (news models.RSSNews, err error) {
	const query1 = `
		select "lastUpdate" from "public.dict" where "idSource" = $1 and "idUser" = $2;
	`
	const query2 = `
		update "public.dict"
		set "lastUpdate" = $1
		where "idSource" = $2 and "idUser" = $3;
	`
	var lStamp time.Time
	err = r.pool.QueryRow(ctx, query1, data.Source, data.ChatID).Scan(&lStamp)
	if err != nil {
		err = app.NotFoundErr
		return
	}

	tmp, lastTS, err := GetRSSNews(data.Source, lStamp)
	if err != nil {
		news = models.RSSNews{News: tmp}
		return
	}
	if tmp != nil {
		_, err = r.pool.Exec(ctx, query2, lastTS, data.Source, data.ChatID)
		news = models.RSSNews{News: tmp}
	} else {
		err = app.HavntUpd
		news = models.RSSNews{News: []string{}}
	}

	return
}

func GetRSSNews(source string, stamp time.Time) ([]string, time.Time, error) {
	var News []string
	var Time time.Time
	var lastTime time.Time
	ext := filepath.Ext(source)

	resp, err := rss.Read(source)
	if err != nil {
		err = app.WrongLink
		return nil, stamp, err
	}

	if ext == ".atom" {
		feed, err := rss.Atom(resp)
		if err != nil {
			fmt.Println(err)
		}

		for _, entry := range feed.Entry {
			println(entry.Updated + " : " + entry.Title)
			Time, err = rss.Date(entry.Updated).Parse()
			if err != nil {
				fmt.Println(err)
			}

			if Time.Unix() > stamp.Unix() {
				News = append(News, fmt.Sprint(entry.Updated+" : "+entry.Title))
				if lastTime.Unix() < Time.Unix() {
					lastTime = Time
				}
			}
		}
	} else {
		channel, err := rss.Regular(resp)
		if err != nil {
			fmt.Println(err)
		}

		for _, item := range channel.Item {
			Time, err = item.PubDate.Parse()
			if err != nil {
				fmt.Println(err)
			}

			if Time.Unix() > stamp.Unix() {
				News = append(News, fmt.Sprint(item.Title+" : "+item.FullText))
				if lastTime.Unix() < Time.Unix() {
					lastTime = Time
				}
			}
		}
	}
	return News, lastTime, nil
}
