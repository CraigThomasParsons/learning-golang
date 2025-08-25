package main

import (
	"log"
	"time"
)

// The struct indexes are capitalized because
// Go is not an object oriented language
// This helps use these things in another package easily
// Capitized variables are public and available in other packages.
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
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "1 (555) 444-1212",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
	log.Println("says")
	firstName, lastName = saySomething(user)

	log.Println(firstName, lastName)
}
