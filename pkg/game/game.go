package game

import (
	"math/rand"
	"slices"
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
