package entity

type Book struct {
	id int
	title string
	author string
	category string
}

func NewBook(id int, title, author, category string) *Book {
	return &Book{
		id: id,
		title: title,
		author: author,
		category: category,
	}
}
