package entity

import "time"

type BorrowedBooks struct {
	book *Book
	date time.Time
}

func NewBorrowedBooks(book *Book, date time.Time) *BorrowedBooks {
	return &BorrowedBooks {
		book: book,
		date: date,
	}
}