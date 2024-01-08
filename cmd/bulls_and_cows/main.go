package main

import (
	"fmt"
	"github.com/dusk-chancellor/bulls_and_cows/internal/game"
)

func main() {
	//LETS GOO
	fmt.Println("Welcome to Bulls and Cows Game!")
	genNum := game.RandomNumberGenerator()
	fmt.Println("Enter your first number: ")
	for i := 1; i <= 10; i++ {
		ok, err := game.PlayGame(i, genNum)
		if err != nil {
			fmt.Printf("Try again !\n")
			i--
		}
		if ok {
			break
		}
	}
}
