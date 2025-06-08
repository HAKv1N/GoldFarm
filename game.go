package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func game() {
	for {
		rand.Seed(time.Now().UnixNano())
		var randomNumber = rand.Intn(numberRange) + 1
		var userInputNumber string
		fmt.Print("Guess the number from 1 to ", numberRange, ": ")
		fmt.Scanln(&userInputNumber)
		userNumber, _ := strconv.Atoi(userInputNumber)

		if randomNumber == userNumber {
			money += moneyAdd
			fmt.Println("Added:", moneyAdd, "money!")
			break
		}
	}
}
