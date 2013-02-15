/*
A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1
Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// From wiki:
	//    A fraction in lowest terms with a prime denominator other than 2 or 5 (i.e. coprime to 10) always produces a repeating decimal.
	//    The period of the repeating decimal of 1/p is equal to the order of 10 modulo p	

	var d int64
	var maxPeriod int = 0

	// skipping 2, 3, 5
	for _, p := range primesUpTo(1000)[3:] {
		period := order(10, p)
		if period > maxPeriod {
			maxPeriod, d = period, p
		}
	}

	fmt.Println(d)
}

// In number theory, given an integer a and a positive integer n with gcd(a,n) = 1, 
// the multiplicative order of a modulo n is the smallest positive integer k with
// a^k = 1 (mod n)
func order(a, n int64) int {
	k := 1
	z, base, mod := big.NewInt(1), big.NewInt(a), big.NewInt(n)

	for {
		// z = base ** k (mod n)
		z.Exp(base, big.NewInt(int64(k)), mod)

		if z.Int64() == 1 {
			break
		}

		k++
	}

	return k
}

func primesUpTo(n int64) []int64 {
	ints := make([]bool, n)

	primes := make([]int64, 0, n/3)
	primes = append(primes, 2)

	var i, j int64

	for i = 3; i < n; i += 2 {
		if !ints[i] {
			primes = append(primes, i)
			for j = i; j < n; j += i {
				ints[j] = true
			}
		}
	}

	return primes
}
