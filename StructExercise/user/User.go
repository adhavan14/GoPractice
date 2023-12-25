package user

import (
	"errors"
	"fmt"
)


type User struct {
	firstname string
	lastname string
	age int
}


func New(firstname, lastname string, age int) (*User, error) {

	if (firstname == "" || lastname == "") {
		return nil, errors.New("firstname and lastname shoul not be empty")
	}
	if (age <= 0) {
		return nil, errors.New("Age should be positive")
	}
	return &User {
		firstname: firstname,
		lastname: lastname,
		age: age,
	}, nil
}

func (user *User) ShowDetails() {
	fmt.Println(user.firstname + "  " + user.lastname + "  " ,user.age)
}

func (user *User) ClearDetails() {
	user.firstname = ""
	user.lastname = ""
	user.age = 0
}