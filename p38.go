/*
Take the number 192 and multiply it by each of 1, 2, and 3:

192 x 1 = 192
192 x 2 = 384
192 x 3 = 576
By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?
*/

package main

import (
	"fmt"
	"strconv"
)

var present struct{} = struct{}{}

func main() {
	maxProduct := 0

	for i := 9000; i < 10000; i++ {
		if product := pandigitalProduct(i); product > maxProduct {
			maxProduct = product
			fmt.Println(maxProduct)
		}
	}
}

func pandigitalProduct(m int) int {
	identity := strconv.Itoa(m)

	for n := 2; len(identity) < 9; n++ {
		identity += strconv.Itoa(m * n)
	}

	if len(identity) > 9 || !isPandigital(identity) {
		return 0
	}

	p, _ := strconv.Atoi(identity)
	return p
}

func isPandigital(s string) bool {
	digits := make(map[rune]struct{})

	for _, d := range s {
		// digits must be 1-9 to be pandigital
		if d != '0' {
			digits[d] = present
		}
	}

	return len(digits) == 9
}
