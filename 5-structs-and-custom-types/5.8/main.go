package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Tyler-Liquornik/golang-fundamentals/5-structs-and-custom-types/5.8/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note instance:", err)
		return
	}

	noteString := userNote.String()
	fmt.Println(noteString)

	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Println("Note saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter note title: ")
	content := getUserInput("Enter note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	// Use bufio to read the entire line of input including spaces,
	// unlike fmt.Scan/Scanln which only read whitespace-separated tokens.
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input, please try again.")
		return getUserInput(prompt)
	}

	// Remove newline characters added when the user presses Enter:
	// "\n" for Unix/macOS, "\r\n" for Windows
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
