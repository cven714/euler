/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	maxPalindrome := 0

	for i := 999; i > 99; i-- {
		for j := i; j > 99; j-- {
			p := i * j

			if isPalindrome(p) && p > maxPalindrome {
				maxPalindrome = p
			}
		}
	}

	fmt.Println(maxPalindrome)
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)

	for start, end := 0, len(s)-1; start <= end; start, end = start+1, end-1 {
		if s[start] != s[end] {
			return false
		}
	}

	return true
}
