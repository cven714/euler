/*
A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(nthPermutation(999999))
}

func nthPermutation(n int) []int {
	factorials := make([]int, 10)
	permutation := make([]int, 10)

	factorials[0] = 1
	for i := 1; i < 10; i++ {
		factorials[i] = factorials[i-1] * i
	}

	for i := 0; i < 10; i++ {
		permutation[i] = n / factorials[10-1-i]
		n = n % factorials[10-1-i]
	}

	for i := 9; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if permutation[j] <= permutation[i] {
				permutation[i]++
			}
		}
	}

	return permutation
}
