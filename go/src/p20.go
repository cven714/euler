/*
n! means n  (n  1)  ...  3  2  1

For example, 10! = 10  9  ...  3  2  1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	sum := 0
	fact100 := factorial(100).String()

	for _, digit := range fact100 {
		if n, err := strconv.Atoi(string(digit)); err == nil {
			sum += n
		}
	}

	fmt.Println(sum)
}

func factorial(n int64) *big.Int {
	if n <= 1 {
		return big.NewInt(1)
	}

	m := big.NewInt(n)
	return m.Mul(m, factorial(n-1))
}
