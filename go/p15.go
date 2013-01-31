/*
(Lattice Paths)
Starting in the top left corner of a 2x2 grid, there are 6 routes (without backtracking) to the bottom right corner.
How many routes are there through a 20x20 grid?
*/

package main

import (
	"fmt"
)

func main() {
	// Justification:
	// You have to make 40 steps to get from top-left to bottom-right,
	// and you have to choose which of those 20 steps are down.
	fmt.Printf("40 choose 20: 137846528820")
}
