package main

import "fmt"

var money int
var numberRange int = 10
var moneyAdd int = 10
var moneyUpgrade int
var numberUpgrade int

const (
	MAX_MONEY_UPGRADE  int = 13
	MAX_NUMBER_UPGRADE int = 6
)

func shop() {
	for {
		fmt.Println("Balance:", money)
		fmt.Print("Choose: m - money, n - number, e - exit: ")
		var userInput string
		fmt.Scanln(&userInput)

		switch userInput {
		case "m":
			for moneyUpgrade <= MAX_MONEY_UPGRADE {
				var moneyUpgradeCost int = (moneyUpgrade + 1) * 400
				fmt.Print("Buy Coin Upgrades ", (moneyUpgrade + 1), " level for ", moneyUpgradeCost, " money?(y/n): ")
				var userBuy string
				fmt.Scanln(&userBuy)

				if userBuy == "y" && money >= moneyUpgradeCost {
					moneyAdd += 10
					moneyUpgrade++
					money -= moneyUpgradeCost
					continue
				} else if userBuy != "y" {
					fmt.Println("Canceling a purchase!")
					break
				} else {
					fmt.Println("Not enough money!")
					break
				}
			}

			if moneyUpgrade >= MAX_MONEY_UPGRADE {
				fmt.Println("The maximum number of coin upgrades has been purchased!")
			}

		case "n":
			for numberUpgrade <= MAX_NUMBER_UPGRADE {
				var numberUpgradeCost int = (numberUpgrade + 1) * 400
				fmt.Print("Buy number upgrades ", (numberUpgrade + 1), " level for ", numberUpgradeCost, " money?(y/n): ")
				var userBuy string
				fmt.Scanln(&userBuy)

				if userBuy == "y" && money >= numberUpgradeCost {
					numberRange -= 1
					numberUpgrade++
					money -= numberUpgradeCost
					continue
				} else if userBuy != "y" {
					fmt.Println("Canceling a purchase!")
					break
				} else {
					fmt.Println("Not enough money!")
					break
				}
			}

			if numberUpgrade >= MAX_NUMBER_UPGRADE {
				fmt.Println("The maximum number of number upgrades has been purchased!")
			}
		}

		if userInput == "e" {
			break
		}
	}
}
