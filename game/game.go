package game

import (
	"fmt"
	"math/rand"
)

var winCondition = map[string]string{
	"Rock":     "Scissors",
	"Papper":   "Rock",
	"Scissors": "Papper"}

var computerChoiceArr = []string{"Rock", "Papper", "Scissors"}

func PlayTurn(playerChoice *string) {
	computer := computerChoiceArr[rand.Intn(3)]
	fmt.Println("Player: " + *playerChoice)
	fmt.Println("Computer: " + computer)
	if winCondition[*playerChoice] == computer {
		fmt.Println("Player win!")
	} else if *playerChoice == computer {
		fmt.Println("It's a draw.")
	} else {
		fmt.Println("Computer win!")
	}
}
