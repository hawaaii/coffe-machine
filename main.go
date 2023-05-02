package main

import (
	"fmt"
	"strconv"
)

func main() {
	machineContents := map[string]int{"water": 400, "milk": 540, "beans": 120, "money": 550, "cups": 9}
	exit := false
  for {
		switch getUserAction() {
		case "buy":
			buyFromMachine(&machineContents)
		case "fill":
			fillMachine(&machineContents)
		case "take":
			submitMoneyMachine(&machineContents)
		case "remaining":
			printMachineStatus(&machineContents)
		case "exit":
      exit = true
		}
    if exit {break}
	}
}

func getUserAction() string {
	var action string
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	fmt.Scan(&action)
	return action
}

func printMachineStatus(machineContents *map[string]int) {
	machine := *machineContents
	fmt.Println("The coffee machine has:")
	fmt.Println(machine["water"], "ml of water")
	fmt.Println(machine["milk"], "ml of milk")
	fmt.Println(machine["beans"], "g of coffee beans")
	fmt.Println(machine["cups"], "disposable cups")
	fmt.Printf("$%v of money \n", machine["money"])
}

func buyFromMachine(machineContents *map[string]int) {
	machine := *machineContents
	menu := map[int][]int{1: {250, 0, 16, 4}, 2: {350, 75, 20, 7}, 3: {200, 100, 12, 6}} // (espresso, latte, cappuccino)
	var buyOption string
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	fmt.Scan(&buyOption)
  if buyOption == "back" {return}
  
  itemSelection, _ := strconv.Atoi(buyOption)
	if machine["water"] < menu[itemSelection][0] {
		fmt.Println("Sorry, not enough water!")
		return
	}
	if machine["milk"] < menu[itemSelection][1] {
		fmt.Println("Sorry, not enough milk!")
		return
	}
	if machine["beans"] < menu[itemSelection][2] {
		fmt.Println("Sorry, not enough water!")
		return
	}
	if machine["cups"] < 1 {
		fmt.Println("Sorry, not enough cups!")
		return
	}
	machine["water"] -= menu[itemSelection][0]
	machine["milk"] -= menu[itemSelection][1]
	machine["beans"] -= menu[itemSelection][2]
	machine["money"] += menu[itemSelection][3]
	machine["cups"] -= 1
	fmt.Println("I have enough resources, making you a coffee!")
}

func fillMachine(machineContents *map[string]int) {
	machine := *machineContents
	var water, milk, coffee, cups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&water)
	fmt.Println("Write how many ml of milk you want to add: ")
	fmt.Scan(&milk)
	fmt.Println("Write how many grams of coffee beans you want to add: ")
	fmt.Scan(&coffee)
	fmt.Println("Write how many disposable cups you want to add: ")
	fmt.Scan(&cups)
	machine["water"] += water
	machine["milk"] += milk
	machine["beans"] += coffee
	machine["cups"] += cups
}

func submitMoneyMachine(machineContents *map[string]int) {
	machine := *machineContents
	fmt.Printf("I gave you $%v \n", machine["money"])
	machine["money"] = 0
}
