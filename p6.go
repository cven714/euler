/*
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import "fmt"

func main() {
	// n(n + 1) / 2
	sum := (100 * 101) / 2
	squareOfSum := sum * sum

	// n(n + 1)(2n + 1) / 6
	sumOfSquares := (100 * 101 * 201) / 6

	fmt.Println(squareOfSum - sumOfSquares)
}
