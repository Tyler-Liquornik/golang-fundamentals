package todo

import (
	"encoding/json"
	"errors"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("text cannot be empty")
	}

	return Todo{
		text,
	}, nil
}

func (t Todo) String() string {
	return "Todo text: " + t.Text
}

func (t Todo) Save() error {
	fileName := "todo.json"
	filePathWithName := "6-interfaces-and-generics/6.1-6.2/" + fileName

	// Note: json.Marshal() will only marshal the exported fields of the struct,
	// so all fields have been made public.
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filePathWithName, jsonBytes, 0644)
}
