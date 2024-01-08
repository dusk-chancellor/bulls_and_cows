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
		ok := game.PlayGame(i, genNum)
		if ok {
			break
		}
	}
}
