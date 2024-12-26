package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstname string
	lastName  string
	birthdate string
	createdAt time.Time
}

// embedded struct
type Admin struct {
	email    string
	password string
	user     User
}

// factory function (constructor)
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}
	return &User{ //returning the memory address of the struct
		firstname: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string, user *User) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}
	return &Admin{
		email:    email,
		password: password,
		user: User{
			firstname: user.firstname,
			lastName:  user.lastName,
			birthdate: user.birthdate,
			createdAt: time.Now(),
		},
	}, nil
}

func (user User) PrintUserData() { //method of the struct
	fmt.Println(user.firstname, user.lastName, user.birthdate)
}

// mutation methods (pass by reference (memory address), not by value (copy of the value inside the address))
func (user *User) ChangeFirstName(newFirstName string) {
	user.firstname = newFirstName
}

func (admin *Admin) PrintSuccessMessage() {
	fmt.Println("Admin created successfully:", admin.email)
}
