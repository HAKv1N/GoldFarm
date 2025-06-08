package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to GoldFarm!")
	fmt.Println("Status:", isAuth)

	auth()
	menu()
}
