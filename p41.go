/*
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?
*/

package main

import (
	"fmt"
	"strconv"

	"utils"
)

var primes []uint64
var present struct{} = struct{}{}

func init() {
	// Can skip 8/9 digit numbers as 
	// 1 + .. + 9 = 45 => div by 3 
	// 1 + .. + 8 = 35 => div by 3
	primes = utils.PrimesUpTo(7654321)
}

func main() {
	for i := len(primes) - 1; i >= 0; i-- {
		if isPandigital(strconv.Itoa(int(primes[i]))) {
			fmt.Println(primes[i])
			break
		}
	}
}

func isPandigital(s string) bool {
	digits := make(map[rune]struct{})
	maxDigit := rune(strconv.Itoa(len(s))[0])

	for _, d := range s {
		// digits must be 1-9 to be pandigital
		if d > '0' && d <= maxDigit {
			digits[d] = present
		}
	}

	return len(digits) == len(s)
}
