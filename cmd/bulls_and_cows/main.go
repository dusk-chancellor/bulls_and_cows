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
	for i := 1; i <= 10; i++ {
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
			default:
				fmt.Printf("Bulls: %v | Cows: %v", bulls, cows)
			}
		} else {
			log.Fatal("an error occurred while handling user input")
		}
	}
}
