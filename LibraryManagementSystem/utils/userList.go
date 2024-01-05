package utils

import (
	"errors"
	"example/App/LibraryManagementSystem/entity"
	"fmt"
	"strings"
)


var listOfAccount []*entity.Account


func CreateAccount() {

	data, err := loadData("resources/users.txt")

	if (err != nil) {
		fmt.Println(err)
		return
	}
	
	for _, d := range data {

		attributes := strings.Split(d, ",")
		account := entity.NewAccount(entity.NewUser(strings.Trim(attributes[0], " "), strings.Trim(attributes[1], " "), strings.Trim(attributes[2], " ")), 
									[]*entity.BorrowedBooks{}, 0.0, true)
		listOfAccount = append(listOfAccount, account)
	}
}

func Authenticate(username, password string) *entity.Account {

	for i := 0; i < len(listOfAccount); i++ {
		if listOfAccount[i].User.AuthenticateUser(username, password) {
			return listOfAccount[i]
		}
	}
	return nil
}

func isUserAvailable(user *entity.User) bool {
	for i := 0; i < len(listOfAccount); i++ {
		if (listOfAccount[i].User).Equals(user) {
			return true
		}
	}
	return false
}


func Register(user *entity.User) (bool, error) {

	if (isUserAvailable(user)) {
		return false, errors.New("Username already exists")
	}

	account := entity.NewAccount(user, []*entity.BorrowedBooks{}, 0.0, true)
	listOfAccount = append(listOfAccount, account)
	return true, nil
}