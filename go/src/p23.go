/*
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/

package main

import (
	"fmt"
)

var abundantNumbers map[int]struct{} = make(map[int]struct{})

func main() {
	sum := 0

	for n := 1; n < 28124; n++ {
		if !sumOfAbundants(n) {
			sum += n
		}

		if isAbundant(n) {
			abundantNumbers[n] = struct{}{}
		}
	}

	fmt.Println(sum)
}

func isAbundant(n int) bool {
	sum := 1

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return sum > n
}

func sumOfAbundants(n int) bool {
	for a, _ := range abundantNumbers {
		// Is n - a also an abundant number?
		// Then n can be written as the sum of two abundant numbers
		if _, ok := abundantNumbers[n-a]; ok {
			return true
		}
	}

	return false
}
