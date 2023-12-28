package main

import (
	"bufio"
	"fmt"
	"github.com/dusk-chancellor/bulls_and_cows/pkg/game"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to Bulls and Cows Game!")
	genNum := game.RandomNumberGeneration()
	fmt.Println("Enter your first number: ")
	for i := 1; i <= 10; i++ {
		fmt.Printf("Attempt №%v\n", i)
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			userInput := scanner.Text()
			bulls, cows, err := game.BullsAndCows(genNum, userInput)
			if err != nil {
				log.Fatal(err)
			}
			switch {
			case bulls == 4:
				fmt.Printf("Congratulations, you won! The number was %v", genNum)
				return
			case i+1 == 11:
				fmt.Printf("You lost. The number was %v", genNum)
			default:
				fmt.Printf("Bulls: %v | Cows: %v\n", bulls, cows)
			}
		} else {
			log.Fatal("an error occurred while handling user input")
		}
	}
}
