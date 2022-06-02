package rss

import (
	"net/http"
	"time"
)

const dateFormat = "Mon, 02 Jan 2006 15:04:05 -0700"

type Fetcher interface {
	Get(url string) (resp *http.Response, err error)
}

type Date string

func (d Date) ParseWithFormat(format string) (time.Time, error) {
	return time.Parse(format, string(d))
}

func (d Date) Parse() (time.Time, error) {
	t, err := d.ParseWithFormat(dateFormat)
	if err != nil {
		t, err = d.ParseWithFormat(time.RFC822) //RSS 2
		if err != nil {
			t, err = d.ParseWithFormat(time.RFC3339) //Atom
		}
	}

	return t, err
}

func Read(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
