/*
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

package main

import (
	"fmt"
	"math/big"
	"strings"
	"strconv"
)

func main() {
	n := big.NewInt(0)
	two := big.NewInt(2)
	exp := big.NewInt(1000)
	
	n = n.Exp(two, exp, nil)
	
	sum := 0	
	for _, digit := range strings.Split(n.String(), "") {
		d, _ := strconv.Atoi(digit)
		sum += d
	}
	
	fmt.Println(sum)
}