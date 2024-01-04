package entity

import "fmt"


type Book struct {
	Id int
	title string
	author string
	category string
}

func NewBook(id int, title, author, category string) *Book {
	return &Book{
		Id: id,
		title: title,
		author: author,
		category: category,
	}
}

func (book *Book) ToStringInBook() string{
	return "Id : " + fmt.Sprintf("%d",book.Id) + " Title : " + book.title + " Author : " + book.author + " Category : " + book.category
}
