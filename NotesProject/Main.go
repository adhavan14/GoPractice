package main

import (
	"bufio"
	"example/App/NotesProject/notes"
	"example/App/NotesProject/JsonOperation"
	"fmt"
	"os"
)


func main() {

	var title, content string
	title, content = getNotesInput()

	var userNotes *notes.Notes
	userNotes, err := notes.New(title, content)

	if (err != nil) {
		fmt.Println(err)
		return
	}
	err = jsonoperation.WriteAsJson("result.json", userNotes)

	if (err != nil) {
		fmt.Println("Saving to file is failed")
		return
	}

	fmt.Println("Saving to file is successful")
}

func getNotesInput() (string, string) {

	title := getUserInput("Enter the title = ")
	content := getUserInput("Enter the content = ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}