package notes

import "errors"
import "fmt"

type Notes struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

func New(title, content string) (*Notes, error) {

	if (len(title) <= 0 || len(content) <=0 ) {
		return nil, errors.New("Title and content is required")
	}

	return &Notes {
		Title: title,
		Content: content,
	}, nil
}

func (notes *Notes) PrintDetails() {
	fmt.Println("Title : " + notes.Title + ", Content : " + notes.Content)
}