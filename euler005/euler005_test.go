package euler005_test

import (
	"github.com/geek4good/kata/euler005"
	"testing"
)

func TestSolution(t *testing.T) {
	actual := euler005.Solution()
	expected := 232792560

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
