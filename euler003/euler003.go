package euler003

import (
	"math"
)

func Problem() string {
	return `
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
`
}

func Solution() int {
	product := 600851475143
	primeFactors := primeFactors(product)
	return primeFactors[len(primeFactors)-1]
}

func primeFactors(product int) []int {
	factors := []int{}
	primes := []int{}
	for i := 2; i <= product; i++ {
		if isPrime(i, &primes) {
			primes = append(primes, i)
			for _, prime := range primes {
				for product%prime == 0 {
					factors = append(factors, prime)
					product /= prime
					if product == 1 {
						return factors
					}
				}
			}
		}
	}
	panic("Failed to find prime factors")
}

func isPrime(number int, primes *[]int) bool {
	for _, prime := range *primes {
		if prime > int(math.Sqrt(float64(number))) {
			break
		}
		if number%prime == 0 {
			return false
		}
	}
	*primes = append(*primes, number)
	return true
}
