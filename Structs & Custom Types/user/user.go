package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	isAdmin  bool
	User
}

func New(firstName, lastName, birthday string) (*User, error) {

	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("first name, last name and birthday are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(firstName, lastName, birthday, email, password string) (*Admin, error) {

	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("first name, last name , birthday , email and password are required")
	}
	return &Admin{
		email:    email,
		password: password,
		isAdmin:  true,
		User: User{
			firstName: firstName,
			lastName:  lastName,
			birthday:  birthday,
			createdAt: time.Now(),
		},
	}, nil
}

func (u *User) printBase() {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name :", u.lastName)
	fmt.Println("Birthday  :", u.birthday)
	fmt.Println("Created At:", u.createdAt.Format(time.DateTime))
}

func (u *User) Print() {
	fmt.Println("User Info")
	fmt.Println("---------")
	u.printBase()
}

func (a *Admin) Print() {
	fmt.Println("Admin Info")
	fmt.Println("----------")
	fmt.Println("Email    :", a.email)
	fmt.Println("Is Admin :", a.isAdmin)
	a.printBase()
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
