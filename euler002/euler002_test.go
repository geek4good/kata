package euler002_test

import (
	"github.com/geek4good/kata/euler002"
	"testing"
)

func TestSolution(t *testing.T) {
	actual := euler002.Solution()
	expected := 4613732

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
