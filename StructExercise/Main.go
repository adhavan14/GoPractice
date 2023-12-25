package main

import (
	"example/App/StructExercise/user"
	"fmt"
)


func main() {
	var appUser *user.User

	appUser, err := user.New("Adhavan", "G", 22)

	if (err != nil) {
		fmt.Println(err)
		return;
	}
	appUser.ShowDetails()
	appUser.ClearDetails()
	appUser.ShowDetails()

	var adminUser *user.Admin

	adminUser, err = user.NewAdmin("adhavan@gmail.com", "adhav@1405")

	if (err != nil) {
		fmt.Println(err)
		return
	}
	adminUser.ShowAdminDetails()
}