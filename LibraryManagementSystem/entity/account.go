package entity

import (
	"errors"
	"fmt"
	"time"
)

const maximumDaysForBorrowBooks int = 5
const finePerDay float64 = 5.0

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
	fine := calculateFine(account, index)
	if fine > 0.0 {
		account.fineAmount = fine
		account.isPaid = false
	}
	account.borrowedBooks = append(account.borrowedBooks[:index], account.borrowedBooks[index+1:]...)
	return nil
}

func (account *Account) IsFinePaid() bool {
	return account.isPaid
}

func (account *Account) PayFine(amount float64) error {

	if (amount > account.fineAmount) {
		return errors.New("Please enter the correct fine amount")
	}
	account.fineAmount = account.fineAmount - amount
	if account.fineAmount == 0.0 {
		account.isPaid = true
	}
	return nil
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

func calculateFine(account *Account, index int) float64 {
	days := int(time.Now().Sub(account.borrowedBooks[index].date).Hours()/24)
	if days > maximumDaysForBorrowBooks {
		return float64(days - maximumDaysForBorrowBooks) * finePerDay
	}
	return 0.0
}