package euler005

import ()

func Problem() string {
	return `
2520 is the smallest number that can be divided by each of the
numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by
all of the numbers from 1 to 20?
`
}

func Solution() int {
	x := 0
	for i := 2520; x == 0; i += 20 {
		for j := 19; j > 2; j-- {
			if i%j != 0 {
				break
			}

			if j == 3 {
				x = i
			}
		}
	}
	return x
}
