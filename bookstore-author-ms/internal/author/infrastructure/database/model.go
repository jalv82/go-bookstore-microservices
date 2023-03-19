package database

const authors = "authors"

type Author struct {
	Id     string `gorm:"primarykey"`
	BookId string
	Name   string
}

func (a *Author) TableName() string {
	return authors
}
