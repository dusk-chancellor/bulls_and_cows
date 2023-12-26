package tests

import (
	"bulls_and_cows/pkg/game"
	"strconv"
	"testing"
)

func TestRandomNumberGenerator(t *testing.T) {
	thefunc := strconv.Itoa(game.RandomNumberGeneration())
	if len(thefunc) != 4 {
		t.Fatalf("The length of number must be 4, not %v", len(thefunc))
	}
}
