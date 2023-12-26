package jsonoperation

import (
	"encoding/json"
	"example/App/NotesProject/notes"
	"os"
	"strings"
)

func WriteAsJson(fileName string, notes *notes.Notes) error {

	fileName = strings.ReplaceAll(fileName, " ", "_")
	jsonNotes, err := json.Marshal(notes)

	if (err != nil) {
		return err
	}
	return os.WriteFile(fileName, jsonNotes, 0644)
}