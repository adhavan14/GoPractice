package utils

import (
	"errors"
	"example/App/LibraryManagementSystem/entity"
)


var listOfAccount []*entity.Account

func CreateAccount() {

	listOfAccount = []*entity.Account{
		entity.NewAccount(entity.NewUser("adhavan@gamil.com", "Adhav@1405", "Adhavan"), []*entity.BorrowedBooks{}, 0.0, true),
		entity.NewAccount(entity.NewUser("meena@gmail.com", "Meena@0405", "Meena"), []*entity.BorrowedBooks{}, 0.0, true),
	}
}


func Authenticate(username, password string) bool {

	for i := 0; i < len(listOfAccount); i++ {
		if listOfAccount[i].User.AuthenticateUser(username, password) {
			return true
		}
	}
	return false
}

func isUserAvailable(username string) bool {

	for i := 0; i < len(listOfAccount); i++ {
		if (listOfAccount[i].User).Equals(username) {
			return true
		}
	}
	return false
}


func Register(username, password, name string) (bool, error) {

	if (isUserAvailable(username)) {
		return false, errors.New("Username already exists")
	}

	account := entity.NewAccount(entity.NewUser(username, password, name), []*entity.BorrowedBooks{}, 0.0, true)
	listOfAccount = append(listOfAccount, account)
	return true, nil
}