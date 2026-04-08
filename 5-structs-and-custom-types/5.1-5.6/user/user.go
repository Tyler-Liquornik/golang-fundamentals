package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if (firstName == "") || (lastName == "") || (birthDate == "") {
		return nil, errors.New("first name, last name, and birth date are required")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email,
		password,
		User{
			FirstName: "Admin",
			LastName:  "Admin",
			BirthDate: "01/01/1970",
			CreatedAt: time.Now(),
		},
	}
}

func (u User) OutputUser() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}
