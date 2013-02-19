/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n := int64(600851475143)
	max := math.Ceil(math.Sqrt(float64(n)))

	out := generate(&max)

	for p := range out {
		if n%p == 0 {
			fmt.Println(p)
		}

		newout := make(chan int64)
		go sieve(p, out, newout)
		out = newout
	}
}

func generate(max *float64) <-chan int64 {
	ch := make(chan int64)

	go func() {
		for n := float64(3); ; n += 2 {
			if n < *max {
				ch <- int64(n)
			} else {
				break
			}
		}

		close(ch)
	}()

	return ch
}

func sieve(prime int64, in <-chan int64, out chan<- int64) {
	for {
		i := <-in

		if i%prime != 0 {
			out <- i
		}
	}
}
