package main

import (
	"fmt"
	"strings"
)

var isExit = false

func menu() {
	for !isExit {
		fmt.Print("Choose action: g - game, s - shop, e - exit: ")
		playerAction := ""

		fmt.Scanln(&playerAction)
		playerAction = strings.ToLower(playerAction)
		playerAction = strings.ReplaceAll(playerAction, " ", "")

		switch playerAction {
		case "g":
			game()
		case "s":
			shop()
		case "e":
			isExit = true
		default:
			fmt.Println("Incorrect action!")
		}
	}
}
