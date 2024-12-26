package main

import (
	"analab/golang/user"
	"fmt"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	newUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	newUser.PrintUserData()
	newUser.ChangeFirstName("Hello")
	newUser.PrintUserData()

	email := getUserData("Please enter your email: ")
	password := getUserData("Please enter your password: ")

	newAdmin, err := user.NewAdmin(email, password, newUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	newAdmin.PrintSuccessMessage()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
