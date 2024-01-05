package utils

import (
	"errors"
	"example/App/LibraryManagementSystem/entity"
	"fmt"
	"strings"
)


var listOfBooks []*entity.Book


func BookList() {

	data, err := loadData("resources/books.txt")

	if (err != nil) {
		fmt.Println(err)
		return
	}
	for index, d := range data {
		attributes := strings.Split(d, ",")
		book := entity.NewBook(index+1, strings.Trim(attributes[0], " "), strings.Trim(attributes[1], " "), strings.Trim(attributes[2], " "))
		listOfBooks = append(listOfBooks, book)
	}
}


func isBookAvailable(book *entity.Book) (bool, int) {

	for i:=0; i<len(listOfBooks); i++ {
		if (book == listOfBooks[i]) {
			return true, i
		}
	}
	return false, -1
}

func List() []*entity.Book{
	return listOfBooks
}

func Add(book *entity.Book) (bool, error){

	isAvailable, _ := isBookAvailable(book)
	if (isAvailable) {
		return false, errors.New("Element already exists")
	}
	listOfBooks = append(listOfBooks, book)
	return true, nil
}

func Delete(book *entity.Book) (bool, error) {
	
	isAvailable, index := isBookAvailable(book)
	if (!isAvailable) {
		return false, errors.New("Element does not exists")
	}

	listOfBooks = append(listOfBooks[:index], listOfBooks[index+1:]...)
	return true, nil
}

func FindById(id int) *entity.Book {

	for i := 0; i < len(listOfBooks); i++ {
		if listOfBooks[i].Id == id {
			return listOfBooks[i]
		}
	}
	return nil
}
