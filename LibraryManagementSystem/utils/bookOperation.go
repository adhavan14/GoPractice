package utils

import "fmt"

func BookOperation() {

	var option string

	fmt.Println(`
1--->View books
2--->Borrow book
3--->Drop book
4--->Pay fine`)
	
	fmt.Print("Choose the option : ")
	fmt.Scan(&option)
}