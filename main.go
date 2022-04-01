package main

import (
	"fmt"
	"rps/game"
)

func main() {
	rounds := 0
	playerChoice := ""
	fmt.Println("Rock Papper Scissors")
	fmt.Printf("====================\n\n\n")
	fmt.Print("Enter amount of rounds: ")
	fmt.Scanln(&rounds)
	fmt.Println()
	for i := 0; i < rounds; i++ {
		fmt.Println("Choose an option: Rock, Papper, Scissors")
		fmt.Print("Enter you'r choice: ")
		fmt.Scanln(&playerChoice)
		game.PlayTurn(&playerChoice)
	}
}
