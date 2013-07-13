/*
A Pythagorean triplet is a set of three natural numbers, a  b  c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import "fmt"

func main() {
	for m := 30; m > 1; m-- {
		for n := m - 1; n > 0; n-- {
			a, b, c := (m*m)-(n*n), 2*m*n, (m*m)+(n*n)
			if a+b+c == 1000 {
				fmt.Printf("%d %d %d -> ", a, b, c)
				fmt.Println(a * b * c)
				break
			}
		}
	}
}
