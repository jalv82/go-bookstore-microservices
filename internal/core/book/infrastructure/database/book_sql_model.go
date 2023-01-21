package database

const books = "books"

type Book struct {
	Id       string `gorm:"primarykey"`
	AuthorId string
	Title    string
}

func (b Book) TableName() string {
	return books
}
