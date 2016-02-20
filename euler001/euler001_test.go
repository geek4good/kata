package euler001_test

import (
	"github.com/geek4good/kata/euler001"
	"testing"
)

func TestSolution(t *testing.T) {
	actual := euler001.Solution()
	expected := 233168

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
