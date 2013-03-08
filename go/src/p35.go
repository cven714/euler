/*
The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
*/

package main

import (
	"fmt"
	"strconv"

	"utils"
)

var primes map[uint64]struct{} = utils.MapPrimesUpTo(1000000)
var circular map[uint64]struct{} = make(map[uint64]struct{})

func main() {
	count := 0

	for prime := range primes {
		if circularPrime(prime) {
			count++
		}
	}

	fmt.Println(count)
}

func circularPrime(p uint64) bool {
	s := strconv.Itoa(int(p))

	for i := 0; i < len(s); i++ {
		rotation, _ := strconv.Atoi(string(s[i:]) + string(s[:i]))
		if _, prime := primes[uint64(rotation)]; !prime {
			return false
		}
	}

	return true
}
