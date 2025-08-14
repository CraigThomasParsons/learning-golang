package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func saySomething(user User) (string, string) {
	log.Println("Hello, from the saySomething func this is", user.FirstName, user.LastName)
	return user.FirstName, user.LastName
}

func main() {
	var firstName string
	var lastName string
	user := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	log.Println(user.FirstName, user.LastName)
	log.Println("says")
	firstName, lastName = saySomething(user)

	log.Println(firstName, lastName)
}
