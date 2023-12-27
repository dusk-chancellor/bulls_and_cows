package game

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

func RandomNumberGeneration() (num string) {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	l := 10
	for i := 0; i <= 3; i++ {
		index := rand.Intn(l)
		n := digits[index]
		num += n
		slices.Delete(digits, index, index+1)
		l--
	}
	return
}

func validCheck(num string) bool {
	if len(num) != 4 {
		return false
	}
	if num[0] == num[1] || num[0] == num[2] || num[0] == num[3] ||
		num[1] == num[2] || num[1] == num[3] || num[2] == num[3] {
		return false
	}
	return true
}

func BullsAndCows(genNum, userInput string) string {
	if !validCheck(userInput) {
		return "ERROR. It seems like your number is not 4-digit or contains at least 1 repeated digit"
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
	res := fmt.Sprintf("Bulls: %v | Cows: %v", bulls, cows)
	return res
}
