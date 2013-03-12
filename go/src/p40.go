/*
An irrational decimal fraction is created by concatenating the positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1 x d10 x d100 x d1000 x d10000 x d100000 x d1000000
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	product, dn := 1, 1
	digit := 0

	for i := 1; i <= 1000000; i++ {
		s := strconv.Itoa(i)
		numLen := len(s)
		digit += numLen

		if digit >= dn {
			c := s[numLen-(digit-dn)-1]
			d, _ := strconv.Atoi(string(c))
			product *= d

			dn *= 10
		}
	}

	fmt.Println(product)
}
