package game

import (
	"math/rand"
)

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func RandomNumberGeneration() (num int) {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	l := 9
	for i := 1; i <= 1000; i *= 10 {
		pick := digits[rand.Intn(l)]
		num += pick * i
		remove(digits, pick)
		l--
	}
	return
}
