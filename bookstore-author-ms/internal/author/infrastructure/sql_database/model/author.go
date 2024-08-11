package model

const Authors = "authors"

type Author struct {
	Id     string `gorm:"primarykey"`
	BookId string
	Name   string
}
