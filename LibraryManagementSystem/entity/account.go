package entity

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
