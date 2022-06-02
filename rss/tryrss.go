package rss

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
)

func Try(stringURL string) string {
	//stringURL := "http://blog.golang.org/feed.atom"
	builder := strings.Builder{}
	ext := filepath.Ext(stringURL)

	u, err := url.Parse(stringURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n" + u.Host)

	resp, err := Read(stringURL)
	if err != nil {
		fmt.Println(err)
	}

	if ext == ".atom" {
		feed, err := Atom(resp)
		if err != nil {
			fmt.Println(err)
		}

		for _, entry := range feed.Entry {
			builder.WriteString(entry.Updated + " : " + entry.Title)
			//fmt.Println(entry.Updated + " " + entry.Title)
		}
	} else {
		channel, err := Regular(resp)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(channel.Title)

		for _, item := range channel.Item {
			//time, err := item.PubDate.Parse()
			if err != nil {
				fmt.Println(err)
			}
			builder.WriteString(item.Title + " : " + item.FullText)
			//fmt.Println(item.Title + " " + item.FullText)
		}
	}
	return builder.String()
}
