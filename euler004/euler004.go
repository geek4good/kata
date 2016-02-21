package euler004

import (
	"sort"
	"strconv"
)

func Problem() string {
	return `
A palindromic number reads the same both ways. The largest
palindrome made from the product of two 2-digit numbers is
9009 = 91 × 99.

Find the largest palindrome made from the product of two
3-digit numbers.
`
}

func Solution() int {
	palindromes := []int{}
	for x := 999; x > 99; x-- {
		for y := x; y > 99; y-- {
			product := x * y
			productString := strconv.Itoa(product)
			if productString == reverse(productString) {
				palindromes = append(palindromes, product)
			}
		}
	}

	sort.Ints(palindromes)
	return palindromes[len(palindromes)-1]
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
