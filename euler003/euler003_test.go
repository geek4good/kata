package euler003_test

import (
	"github.com/geek4good/kata/euler003"
	"testing"
)

func TestSolution(t *testing.T) {
	actual := euler003.Solution()
	expected := 6857

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
