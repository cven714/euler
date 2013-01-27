/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumNPrimes(2000000))
}

func sumNPrimes(n int64) int64 {
	ints := make([]bool, n)
	sum := int64(2)
	
	var i, j int64
	
	for i = 3; i < n; i += 2 {
		if !ints[i] {
			sum += i
			for j = i; j < n; j += i {
				ints[j] = true
			}
		}
	}
	
	return sum
}
