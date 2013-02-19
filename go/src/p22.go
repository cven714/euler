/*
Using names.txt, a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938  53 = 49714.

What is the total of all the name scores in the file?
*/

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	names := sort.StringSlice(readNames())
	names.Sort()

	var sum uint64 = 0
	for rank, name := range names {
		score := 0
		for _, letter := range name {
			// Subtracting 64 so rune A (65) -> 1, B (66) -> 2, etc.
			score += int(letter - 64)
		}

		sum += uint64(score * (rank + 1))
	}

	fmt.Println(sum)
}

func readNames() []string {
	data, _ := ioutil.ReadFile("p22names.txt")
	return strings.Split(string(data), ",")
}
