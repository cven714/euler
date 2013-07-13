/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a  b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

package main

import (
	"fmt"
)

func main() {
	amicable := make(map[int]struct{})
	sum := 0

	for n := 1; n < 10000; n++ {
		if _, ok := amicable[n]; !ok {
			dn := d(n)
			ddn := d(dn)

			if n != dn && n == ddn {
				fmt.Printf("d(%d) = %d, d(%d) = %d \n\n", n, dn, dn, ddn)
				sum += n + dn
			}

			amicable[n] = struct{}{}
			amicable[dn] = struct{}{}
		}
	}

	fmt.Println(sum)
}

func d(n int) int {
	sum := 0

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return sum
}
