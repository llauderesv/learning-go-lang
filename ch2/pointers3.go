package main

import "fmt"

type User struct {
	email string
	name  string
	age   int
}

func (u *User) updateEmail(email string) {
	u.email = email
}

func updateEmail(u *User, email string) {
	u.email = email
}

func main() {
	vince := User{
		email: "Vincent.Llauderes@gmail.com",
		name:  "Vincent Llauderes",
		age:   31,
	}

	fmt.Println("Email before update: ", vince.email)

	// Same with above but. Above code is a more syntactic sugar and ensures that update email is in the context of User.
	// vince.updateEmail("AppleOPangantihon@gmail.com")
	updateEmail(&vince, "AppleOPangantihon@gmail.com")

	fmt.Println("Email after update: ", vince.email)
}
