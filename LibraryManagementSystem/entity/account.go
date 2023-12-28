package entity

type Account struct {
	user User
	borrowedBooks string
	fineAmount float64
	isPaid bool
}