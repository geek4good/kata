package euler1_test

import (
	"github.com/geek4good/kata/euler1"
	"testing"
)

func TestSolution(t *testing.T) {
	actual := euler1.Solution()
	expected := 233168

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
