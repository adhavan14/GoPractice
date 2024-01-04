package utils

import (
	"example/App/LibraryManagementSystem/entity"
	"fmt"
	"time"
)

func BookOperation(account *entity.Account) {

	for {
		var option string

		fmt.Println(`
	1--->View books
	2--->Borrow book
	3--->Drop book
	4--->Pay fine
	5--->Profile
	other--->Exit`)
		
		fmt.Print("Choose the option : ")
		fmt.Scan(&option)

		switch option {
			case "1": totalBooks := List()
					for i := 0; i < len(totalBooks); i++ {
						fmt.Println(totalBooks[i].ToStringInBook())
					}
			case "2": borrowBook(account)
			case "3": dropBook(account)
			case "4": payFine()
			case "5": fmt.Println(account.ToStringInAccount())
			default : return
		}
	}
}

func borrowBook(account *entity.Account) {

	var id int
	fmt.Print("Enter the id of the book : ")
	fmt.Scan(&id)

	book := FindById(id)
	if (book == nil) {
		fmt.Println("Enter the valid id")
		return
	}
	borrowBook := entity.NewBorrowedBooks(book, time.Now())
	err := account.BorrowBook(borrowBook)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println("Book borrowed successfully")
}

func dropBook(account *entity.Account) {
	var id int
	fmt.Print("Enter the id of the book : ")
	fmt.Scan(&id)

	book := FindById(id)
	if (book == nil) {
		fmt.Println("Enter the valid id")
		return
	}

	account.DropBook(book)
}

func payFine() {

}
