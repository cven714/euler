/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import "fmt"
import "io/ioutil"
import "strconv"
import "strings"

func main() {
	sums := make(chan int64)
	go sumFile("prime1-1mil.txt", sums)
	go sumFile("prime1mil-2mil.txt", sums)
	
	sum1 := <- sums
	sum2 := <- sums
	
	fmt.Println(sum1)
	fmt.Println(sum2)
	
	fmt.Println(sum1 + sum2)
}

func sumFile(fileName string, out chan<- int64) {
	file, _ := ioutil.ReadFile(fileName)
	lines := strings.Split(string(file), "\r")
	
	var sum int64 = 0
	
	for _, line := range lines {
		for _, n := range strings.Split(line, ",") {
			prime, _ := strconv.Atoi(n)			
			sum += int64(prime)
		}
	}
	
	out <- sum
}
