/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?


NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/

package main

import (
	"fmt"
)

type number int

var onesNames = map[number]string{
	0: "",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var teensNames = map[number]string{
	0:  "",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var tensNames = map[number]string{
	0: "",
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

func main() {
	// 3 letters taken away 9 times to compensate for 100 -> one hundred 'and'
	letterCount := -3 * 9
	letterCount += len("onethousand")
	for i := 1; i < 1000; i++ {
		fmt.Println(number(i))
		letterCount += len(number(i).String())
	}

	fmt.Println(letterCount)
}

func (n number) String() string {
	hundreds, tens, ones := n/100, (n%100)/10, n%10
	var hundredsText, tensText, onesText string

	if hundreds > 0 {
		hundredsText = fmt.Sprintf("%shundredand", onesNames[hundreds])
	}

	if tens == 1 {
		tensText = teensNames[ones+10]
		ones = 0
	} else {
		tensText = tensNames[tens]
	}

	onesText = onesNames[ones]

	return fmt.Sprintf("%s%s%s", hundredsText, tensText, onesText)
}
