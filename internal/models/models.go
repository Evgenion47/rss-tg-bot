package models

type User struct {
	ChatID int64
}

type Source struct {
	Source string
}

type Sources struct {
	Sources []string
}
type RSSNews struct {
	News []string
}

type DCData struct {
	ChatID int64
	Source string
}
