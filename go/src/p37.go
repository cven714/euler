/*
The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/

package main

import (
	"fmt"
	"strconv"
	"utils"
)

var primes []uint64 = utils.PrimesUpTo(1000000)
var present struct{} = struct{}{}
var primesMap map[uint64]struct{} = map[uint64]struct{}{
	2: present,
	3: present,
	5: present,
	7: present,
}

var truncatable map[uint64]struct{} = make(map[uint64]struct{})

func main() {
	// Start at the first prime > 10
	for i := 4; i < len(primes); i++ {
		p := primes[i]
		primesMap[p] = present

		if truncatablePrime(p) {
			truncatable[p] = present
		}
	}

	var sum uint64 = 0
	for p, _ := range truncatable {
		sum += p
	}

	fmt.Println(sum)
}

func truncatablePrime(p uint64) bool {
	s := strconv.Itoa(int(p))

	for i := 0; i < len(s); i++ {
		left, _ := strconv.Atoi(s[:len(s)-i])
		right, _ := strconv.Atoi(s[i:])

		if !isPrime(left) || !isPrime(right) {
			return false
		}
	}

	return true

}

func isPrime(n int) bool {
	_, prime := primesMap[uint64(n)]

	return prime
}
