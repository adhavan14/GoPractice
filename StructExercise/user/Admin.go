package user

import (
	"errors"
	"fmt"
)


type Admin struct {
	email string
	password string
	user *User
}


func NewAdmin(email, password string) (*Admin, error) {

	appUser, error := New("Admin", "owner", 20)
	if (error != nil || email == "" || password == "") {
		return nil, errors.New("All fields are required")
	}
	return &Admin{
		email: email,
		password: password,
		user: appUser,
	}, nil
}

func (admin *Admin) ShowAdminDetails() {
	fmt.Println(admin.email + " " + admin.password)
	admin.user.ShowDetails()
}