package main

import "fmt"

var isAuth bool
var user User

type User struct {
	name string
	age  int
}

func auth() {
	for !isAuth {
		fmt.Print("You are not logged in! Enter a name: ")
		name := ""
		fmt.Scanln(&name)
		fmt.Print("Enter the age: ")
		age := 0
		fmt.Scanln(&age)

		if name != "" && age > 0 && age < 150 {
			user.name = name
			user.age = age
			isAuth = true
			break
		} else {
			fmt.Println("Incorrect data!")
		}
	}

	if isAuth {
		fmt.Println("Your data:", user)
	}
}
