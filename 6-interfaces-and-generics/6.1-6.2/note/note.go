package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}

	return Note{
		title,
		content,
		time.Now(),
	}, nil
}

func (n Note) String() string {
	return "Note title: " + n.Title + "\nNote content: " + n.Content + "\nNote created at: " + n.CreatedAt.String()
}

func (n Note) Save() error {
	fileName := strings.ToLower(strings.ReplaceAll(n.Title, " ", "_")) + ".json"
	filePathWithName := "6-interfaces-and-generics/6.1-6.2/" + fileName

	// Note: json.Marshal() will only marshal the exported fields of the struct,
	// so all fields have been made public.
	jsonBytes, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(filePathWithName, jsonBytes, 0644)
}
