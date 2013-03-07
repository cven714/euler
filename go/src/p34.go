/*
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/

package main

import (
	"fmt"
)

var factorial = map[int]int{
	0: 1,
	1: 1,
	2: 2,
	3: 6,
	4: 24,
	5: 120,
	6: 720,
	7: 5040,
	8: 40320,
	9: 362880,
}

func main() {
	sum := 0

	for i := 10; i < 2500000; i++ {
		n := i
		digitSum := 0

		for n > 0 {
			digit := n % 10
			digitSum += factorial[digit]
			n /= 10
		}

		if digitSum == i {
			fmt.Println(i)
			sum += i
		}
	}

	fmt.Println(sum)
}
