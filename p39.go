/*
If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p < 1000, is the number of solutions maximised?
*/

package main

import (
	"fmt"
)

var perimeterSolutions map[int]int = make(map[int]int)
var triples map[string]struct{} = make(map[string]struct{})

func main() {
	for m := 2; m < 30; m++ {
		for n := 1; n < m; n++ {
			for k := 0; ; k++ {
				a, b, c := k*((m*m)-(n*n)), k*(2*m*n), k*((m*m)+(n*n))
				p := a + b + c

				if p <= 1000 {
					addSolution(a, b, c, p)
				} else {
					break
				}
			}
		}
	}

	maxP := 0
	maxSolutions := 0

	for p, solutions := range perimeterSolutions {
		if solutions > maxSolutions {
			maxSolutions, maxP = solutions, p
		}
	}

	fmt.Println(maxP)
}

func addSolution(a, b, c, p int) {
	// Ensure a < b < c
	if a > b {
		a, b = b, a
	}

	identity := fmt.Sprintf("%d %d %d", a, b, c)

	if _, ok := triples[identity]; ok {
		return
	}

	if _, ok := perimeterSolutions[p]; !ok {
		perimeterSolutions[p] = 1
	} else {
		perimeterSolutions[p]++
	}

	triples[identity] = struct{}{}
}
