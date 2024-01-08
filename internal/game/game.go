package game

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"slices"
	"strings"
)

func RandomNumberGenerator() string { //4-digit random number generator
	//digits are non-repeated & "0" must be able to stay at the front
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var num string
	for i := 0; i <= 3; i++ {
		index := rand.Intn(len(digits))
		n := digits[index]
		num += n
		slices.Delete(digits, index, index+1)
	}
	return num
}

func validCheck(num string) bool { //checking user input number
	//4-digit only
	if len(num) != 4 {
		return false
	}
	//not any repeated digits
	if num[0] == num[1] || num[0] == num[2] || num[0] == num[3] ||
		num[1] == num[2] || num[1] == num[3] || num[2] == num[3] {
		return false
	}
	return true
}

func BullsAndCows(genNum, userInput string) (int, int, error) {
	//main function of the game - checks if user input number has any bulls and cows
	//e.g. generated number: 4071 | user input: 1024 | -> '0' is bull, '1' and '4' are cows
	if !validCheck(userInput) {
		return 0, 0, errors.New(
			"error while checking number (perhaps digits are repeated or not 4-digit number is entered)")
	}
	var bulls, cows int
	for i, ch := range genNum {
		if genNum[i] == userInput[i] {
			bulls++
		}
		if strings.ContainsAny(userInput, string(ch)) && genNum[i] != userInput[i] {
			cows++
		}
	}
	return bulls, cows, nil
}

func PlayGame(i int, genNum string) bool { //handles game in terminal
	//
	fmt.Printf("Attempt №%v\n", i)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput := scanner.Text()
		bulls, cows, err := BullsAndCows(genNum, userInput)
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case bulls == 4:
			fmt.Printf("Congratulations, you won! The number was %v!\n", genNum)
			return true
		case i+1 == 11:
			fmt.Printf("You lost. The number was %v\n", genNum)
		default:
			fmt.Printf("Bulls: %v | Cows: %v\n", bulls, cows)
		}
	} else {
		log.Fatal("an error occurred while handling user input")
	}
	return false
}
