package main

import (
	"fmt"

	"example.com/user"
)

func main() {
	for {
		var choice int

		fmt.Println("Choose The Type Of User:")
		fmt.Println("1 - Admin")
		fmt.Println("2 - User")
		fmt.Println("3 - Exit")
		fmt.Print("Your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			handleAdmin()
		case 2:
			handleUser()
		case 3:
			fmt.Println("Goodbye 👋")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}

		fmt.Println()
	}
}

func handleUser() {
	var firstName, lastName, birthday string

	fmt.Print("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Enter Your Birthday: ")
	fmt.Scanln(&birthday)

	u := user.New(firstName, lastName, birthday)
	u.Print()
}

func handleAdmin() {
	var firstName, lastName, birthday, email, password string

	fmt.Print("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Print("Enter Your Password: ")
	fmt.Scanln(&password)

	fmt.Print("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Enter Your Birthday: ")
	fmt.Scanln(&birthday)

	a := user.NewAdmin(firstName, lastName, birthday, email, password)
	a.Print()
}
