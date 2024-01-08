package game

import (
	"testing"
)

func TestRandomNumberGenerator(t *testing.T) { //testing RandomNumberGenerator()
	theFunc := RandomNumberGenerator()
	if len(theFunc) != 4 {
		t.Fatalf("The length of number must be 4, not %v", len(theFunc))
	}
	if theFunc[0] == theFunc[1] || theFunc[0] == theFunc[2] || theFunc[0] == theFunc[3] ||
		theFunc[1] == theFunc[2] || theFunc[1] == theFunc[3] || theFunc[2] == theFunc[3] {
		t.Fatalf("Digits within the number are repeated")
	}
}
