package models

type User struct {
	UserID int64
}

type Source struct {
	Source  string
	LastUpd int64
}

type DCData struct {
	UserID int64
	Source string
}
