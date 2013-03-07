/*
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 * 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
*/

package main

import (
	"fmt"
)

var present struct{} = struct{}{}

var products map[int]struct{} = make(map[int]struct{})

func main() {
	for i := 1; i < 2000; i++ {
		for j := 1; j < i; j++ {
			p := i * j
			if pandigital(i, j, p) {
				products[p] = present
			}
		}
	}

	sum := 0
	for p := range products {
		sum += p
	}

	fmt.Println(sum)
}

func pandigital(m, n, p int) bool {
	identity := fmt.Sprintf("%d%d%d", m, n, p)
	if len(identity) != 9 {
		return false
	}

	digits := make(map[rune]struct{})

	for _, d := range identity {
		// digits must be 1-9 to be pandigital
		if d != '0' {
			digits[d] = present
		}
	}

	return len(digits) == 9
}
