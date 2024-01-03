package main

import (
	"example/App/LibraryManagementSystem/entity"
	"example/App/LibraryManagementSystem/utils"
	"fmt"
	"os"
)

func init() {
	utils.BookList()
	utils.CreateAccount()
}

func main() {
	fmt.Println("------------Welcome-----------")
	fmt.Println("Library Management System")

	var option string
	fmt.Println("1---->Register")
	fmt.Println("2---->Login")
	fmt.Println("3---->Exit")
	fmt.Println()
	fmt.Println("Choose the option : ")
	fmt.Scan(&option)

	switch option {
	case "1":
		err := register()
		if (err != nil) {
			fmt.Println(err)
			return
		} 
		fmt.Println("User registered successfully")
		fmt.Println("Please login")
		result := login()
		if (!result) {
			fmt.Println("Login failed")
			return
		}
		fmt.Println("Login successful")
		utils.BookOperation()
	case "2":
		result := login()
		if (!result) {
			fmt.Println("Login failed")
			return
		}
		fmt.Println("Login successful")
		utils.BookOperation()
	case "3":
		os.Exit(0)
	default:
		fmt.Println("Choose the valid option")
	}
}


func login() bool {

	var username, password string
	username = getInput("Enter the username : ")
	password = getInput("Enter the password : ")
	result := utils.Authenticate(username, password)
	return result
}

func register() error {

	var username, password, name string
	username = getInput("Enter the username : ")
	password = getInput("Enter the password : ")
	name = getInput("Enter your name : ")
	user := entity.NewUser(username, password, name)
	_, err := utils.Register(user)
	return err
}

func getInput(message string) string {

	var input string
	fmt.Print(message)
	fmt.Scan(&input)
	return input
}