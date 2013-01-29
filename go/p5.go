/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import "fmt"

func main() {
	// Just need enough prime factors to make all numbers <= 20
	// 2 = 2
	// 3 = 3
	// 4 = 2 * 2
	// 5 = 5
	// 6 = 2 * 3
	// 7 = 7
	// 8 = 2 * 2 * 2
	// 9 = 3 * 3
	// 10 = 2 * 5
	// 11 = 11
	// 12 = 2 * 2 * 3
	// 13 = 13
	// 14 = 2 * 7
	// 15 = 3 * 5
	// 16 = 2 * 2 * 2 * 2
	// 17 = 17
	// 18 = 2 * 3 * 3
	// 19 = 19
	// 20 = 2 * 2 * 5

	largest := 1

	// Multiply each factor the maximum number of times it appears in the prime factorings
	largest *= 2 * 2 * 2 * 2
	largest *= 3 * 3
	largest *= 5
	largest *= 7
	largest *= 11
	largest *= 13
	largest *= 17
	largest *= 19

	fmt.Println(largest)
}
