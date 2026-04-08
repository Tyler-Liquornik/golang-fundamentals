package main

import (
	"fmt"

	"github.com/Tyler-Liquornik/golang-fundamentals/5-structs-and-custom-types/5.1-5.6/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	adminUser := user.NewAdmin("test@example.com", "password")

	appUser.OutputUser()
	adminUser.OutputUser()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
