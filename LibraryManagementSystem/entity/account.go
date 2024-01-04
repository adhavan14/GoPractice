package entity

import (
	"errors"
	"fmt"
)

type Account struct {
	User *User
	borrowedBooks []*BorrowedBooks
	fineAmount float64
	isPaid bool
}

func NewAccount(user *User, borrowedBooks []*BorrowedBooks, fineAmount float64, isPaid bool) *Account {
	return &Account{
		User: user,
		borrowedBooks: borrowedBooks,
		fineAmount: fineAmount,
		isPaid: isPaid,
	}
}

func isAlreadyBorrowed(account *Account,book *Book) (bool, int) {
	for i := 0; i<len(account.borrowedBooks); i++ {
		if account.borrowedBooks[i].book == book {
			return true, i
		}
	}
	return false, -1
}

func (account *Account) BorrowBook(borrowBook *BorrowedBooks) error {

	isBorrowed, _ := isAlreadyBorrowed(account, borrowBook.book)
	if  isBorrowed {
		return errors.New(borrowBook.book.title + " is already borrowed")
	}
	account.borrowedBooks = append(account.borrowedBooks, borrowBook)
	return nil
}

func (account *Account) DropBook(book *Book) error {
	isBorrowed, index := isAlreadyBorrowed(account, book)
	if !isBorrowed {
		return errors.New(book.title + " is not borrowed in your account")
	}
	account.borrowedBooks = append(account.borrowedBooks[:index], account.borrowedBooks[index+1:]...)
	return nil
}

func (account *Account) PayFine() {

}

func (account *Account) ToStringInAccount() string {
	var user string =  account.User.ToStringInUser() + "\n"
	var borrowBook string = "[\n"
	for i := 0; i < len(account.borrowedBooks); i++ {
		borrowBook += account.borrowedBooks[i].ToStringInBorrowedBook() + "\n"
	}
	borrowBook += "]\n"
	return "User : " + user + "Borrowed books : " + borrowBook + "Fine : " + fmt.Sprintf("%0.2f", account.fineAmount)
}