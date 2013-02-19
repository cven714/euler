/*
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in triangle.txt (right click and 'Save Link/Target As...'), a 15K text file containing a triangle with one-hundred rows.

NOTE: This is a much more difficult version of Problem 18. It is not possible to try every route to solve this problem, as there are 299 altogether! If you could check one trillion (1012) routes every second it would take over twenty billion years to check them all. There is an efficient algorithm to solve it. ;o)
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func init() {
	b, _ := ioutil.ReadFile("p67triangle.txt")
	s := strings.Replace(string(b), "\n", " ", -1)
	numData := strings.Split(s, " ")

	for _, n := range numData {
		number, err := strconv.Atoi(n)
		if err == nil {
			data = append(data, number)
		}
	}
}

func main() {
	buildTriangle()
	fmt.Println(getMaxTreeSum(triangle[0]))
}

type tree struct {
	left, right *tree
	value, id   int
}

var data []int
var triangle = make(map[int]*tree)
var maxTreeSum = make(map[int]int)

func getMaxTreeSum(t *tree) int {
	sum := 0
	if t != nil {
		var ok bool
		if sum, ok = maxTreeSum[t.id]; !ok {
			maxLeft := getMaxTreeSum(t.left)
			maxRight := getMaxTreeSum(t.right)

			if maxLeft > maxRight {
				sum = t.value + maxLeft
			} else {
				sum = t.value + maxRight
			}
		}

		maxTreeSum[t.id] = sum
	}

	return sum
}

func buildTriangle() {
	rowLen, remainingInRow := 1, 1
	for i := 0; i < len(data); i, remainingInRow = i+1, remainingInRow-1 {
		if remainingInRow == 0 {
			rowLen++
			remainingInRow = rowLen
		}

		t := getTree(i)
		t.left = getTree(i + rowLen)
		t.right = getTree(i + rowLen + 1)
	}
}

func getTree(index int) *tree {
	t := triangle[index]
	if t == nil && index < len(data) {
		t = &tree{id: index, value: data[index]}
		triangle[index] = t
	}

	return t
}
