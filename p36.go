/*
The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum, i int64 = 0, 0

	for i = 0; i < 1000000; i++ {
		if isPalindrome(strconv.FormatInt(i, 10)) && isPalindrome(strconv.FormatInt(i, 2)) {
			sum += i
		}
	}

	fmt.Println(sum)
}

func isPalindrome(s string) bool {
	for start, end := 0, len(s)-1; start <= end; start, end = start+1, end-1 {
		if s[start] != s[end] {
			return false
		}
	}

	return true
}
