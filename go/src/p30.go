/*
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 1^4 + 6^4 + 3^4 + 4^4
8208 = 8^4 + 2^4 + 0^4 + 8^4
9474 = 9^4 + 4^4 + 7^4 + 4^4
As 1 = 14 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/

package main

import (
	"fmt"
)

func main() {
	sum := 0

	for i := 2; i < (9*9*9*9*9)*6; i++ {
		htou := (i % 1000000) / 100000
		ttou := (i % 100000) / 10000
		thou := (i % 10000) / 1000
		hund := (i % 1000) / 100
		tens := (i % 100) / 10
		ones := (i % 10) / 1

		digitSum :=
			(htou * htou * htou * htou * htou) +
				(ttou * ttou * ttou * ttou * ttou) +
				(thou * thou * thou * thou * thou) +
				(hund * hund * hund * hund * hund) +
				(tens * tens * tens * tens * tens) +
				(ones * ones * ones * ones * ones)

		if digitSum == i {
			sum += i
			fmt.Println(htou, ttou, thou, hund, tens, ones)
			fmt.Println(i)
		}
	}

	fmt.Println(sum)
}
