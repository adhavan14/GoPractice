package utils

import (
	"errors"
	"example/App/LibraryManagementSystem/entity"
)


var listOfBooks []*entity.Book

func bookList() {
	listOfBooks = []*entity.Book {
		entity.NewBook(1, "The Grass is Always Greener", "Jeffrey Archer", "Modern Times"),
		entity.NewBook(2, "Murder!", "Arnold Bennett", "Crime"),
		entity.NewBook(3, "An Occurrence at Owl Creek Bridge One of the Missing", "Ambrose Bierce",	"Adventure"),
		entity.NewBook(4, "The Higgler", "A. E. Coppard", "Romance"),
		entity.NewBook(5, "The Open Boat", "Stephen Crane", "Classic"),
		entity.NewBook(6, "The Speckled Band", "Sir Arthur Conan Doyle", "Crime"),
		entity.NewBook(7, "The Signalman", "Charles Dickens", "Suspense"),
		entity.NewBook(8, "The Five Orange Pips", "Arthur Conan Doyle",	"Crime"),
		entity.NewBook(9, "From a View to a Kill", "Ian Fleming", "Adventure"),
		entity.NewBook(10, "A Mere Interlude", "Thomas Hardy", "Romance"),
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

	listOfBooks = append(listOfBooks[:index], listOfBooks[index +1:]...)
	return true, nil
}

