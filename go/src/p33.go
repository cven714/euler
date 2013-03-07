/*
The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
*/

package main

import (
	"fmt"
)

func main() {
	for d := 11; d < 99; d++ {
		for n := 10; n < d; n++ {
			decimal := float64(n) / float64(d)

			nTens, nOnes := n/10, n%10
			dTens, dOnes := d/10, d%10

			if nOnes == dTens {
				reduced := float64(nTens) / float64(dOnes)
				if decimal == reduced {
					fmt.Printf("%d/%d == %d/%d\n", n, d, nTens, dOnes)
				}
			}
		}
	}
}
