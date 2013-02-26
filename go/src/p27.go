/*
Euler published the remarkable quadratic formula:

n^2 + n + 41

It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39. However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and certainly when n = 41, 41^2 + 41 + 41 is clearly divisible by 41.

Using computers, the incredible formula  n^2 - 79n + 1601 was discovered, which produces 80 primes for the consecutive values n = 0 to 79. The product of the coefficients, 79 and 1601, is 126479.

Considering quadratics of the form:

n^2 + an + b, where |a| < 1000 and |b| < 1000

where |n| is the modulus/absolute value of n
e.g. |11| = 11 and |4| = 4
Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n = 0.
*/

package main

import (
	"fmt"
	"utils"
)

var primes map[int]struct{}

func init() {
	prime := struct{}{}
	primes = make(map[int]struct{})

	for _, p := range utils.PrimesUpTo(10000) {
		primes[int(p)] = prime
	}
}

func main() {
	maxConsecutive, ab := 0, 1

	for a := 0; a < 1000; a++ {
		if _, aPrime := primes[a]; aPrime {
			for b := 0; b < 1000; b++ {
				if _, bPrime := primes[b]; bPrime {
					if consecutive, prod := consecutivePrimes(a, b), a*b; consecutive > maxConsecutive {
						maxConsecutive, ab = consecutive, prod
					}
					if consecutive, prod := consecutivePrimes(-a, b), -a*b; consecutive > maxConsecutive {
						maxConsecutive, ab = consecutive, prod
					}
					if consecutive, prod := consecutivePrimes(a, -b), a*-b; consecutive > maxConsecutive {
						maxConsecutive, ab = consecutive, prod
					}
					if consecutive, prod := consecutivePrimes(-a, -b), -a*-b; consecutive > maxConsecutive {
						maxConsecutive, ab = consecutive, prod
					}
				}
			}
		}
	}

	fmt.Println(maxConsecutive, ab)
}

func consecutivePrimes(a, b int) int {
	n := 0
	for {
		fn := (n * n) + (a * n) + b

		if _, prime := primes[fn]; !prime {
			break
		}

		n++
	}

	return n
}
