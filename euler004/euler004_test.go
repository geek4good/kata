package euler004_test

import (
	"github.com/geek4good/kata/euler004"
	"testing"
)

func TestSolution(t *testing.T) {
	actual := euler004.Solution()
	expected := 906609

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
