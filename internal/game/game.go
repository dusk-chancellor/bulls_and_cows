package game

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type InputErrors struct { //struct for checking user input number
	Msg string
}

func (e *InputErrors) Error() string { //method
	return e.Msg
}

func deleteFunc(slice []string, i int) []string { //deletes fixed element in slice
	//it is done for further implementations
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}

func RandomNumberGenerator() string { //4-digit random number generator
	//digits are non-repeated & "0" must be able to stay at the front
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var num string
	for i := 0; i <= 3; i++ {
		index := rand.Intn(len(digits))
		num += digits[index]
		deleteFunc(digits, index)
	}
	return num
}

func validCheck(num string) error { //checking user input number
	//must be number, nothing else !
	for _, ch := range num {
		if ch > '9' || ch < '0' {
			return &InputErrors{
				Msg: fmt.Sprintf("'%v' is not a number", string(ch)),
			}
		}
	}
	//4-digit only
	if len(num) != 4 {
		return &InputErrors{
			Msg: fmt.Sprintf("'%v' is not %d-digit number, not 4-digit",
				num, len(num)),
		}
	}
	//not any repeated digits
	elements := make(map[string]int)
	for _, n := range num {
		if elements[string(n)] == 1 {
			elements[string(n)]++
		} else {
			elements[string(n)] = 1
		}
	}

	for key, val := range elements {
		if elements[key] > 1 {
			return &InputErrors{
				Msg: fmt.Sprintf("digit '%v' is repeated %d times",
					key, val),
			}
		}
	}
	return nil
}

func BullsAndCows(genNum, userInput string) (int, int, error) {
	//main function of the game - checks if user input number has any bulls and cows
	//e.g. generated number: 4071 | user input: 1024 | -> '0' is bull, '1' and '4' are cows
	err := validCheck(userInput)
	if err != nil {
		var inputError *InputErrors
		if errors.As(err, &inputError) {
			fmt.Printf("'%v' is not valid: %s\n",
				userInput, inputError.Error())
			return 0, 0, err
		}
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

func PlayGame(i int, genNum string) (bool, error) { //handles game in terminal
	//
	fmt.Printf("Attempt №%v\n", i)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput := scanner.Text()
		bulls, cows, err := BullsAndCows(genNum, userInput)
		if err != nil {
			return false, err
		}

		switch {
		case bulls == 4:
			fmt.Printf("Congratulations, you won! The number was %v!\n", genNum)
			return true, nil
		case i == 10:
			fmt.Printf("You lost. The number was %v\n", genNum)
		default:
			fmt.Printf("Bulls: %v | Cows: %v\n", bulls, cows)
		}
	} else {
		log.Fatal("an error occurred while handling user input")
	}
	return false, nil
}
