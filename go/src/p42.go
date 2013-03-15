/*
The nth term of the sequence of triangle numbers is given by, tn =n(n+1)/2; so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var triangle map[int]struct{} = make(map[int]struct{})

func init() {
	for i, t := 1, 1; t < 1000; i, t = i+1, t+i+1 {
		triangle[t] = struct{}{}
	}
}

func main() {
	triangularWords := 0

	data, _ := ioutil.ReadFile("p42words.txt")
	for _, word := range strings.Split(string(data), ",") {
		word = strings.Replace(word, "\"", "", 2)
		if _, triangular := triangle[wordValue(word)]; triangular {
			triangularWords++
		}
	}

	fmt.Println(triangularWords)
}

func wordValue(s string) int {
	value := 0
	for _, l := range s {
		// Subtracting 64 so rune A (65) -> 1, B (66) -> 2, etc.
		value += int(l - 64)
	}

	return value
}
