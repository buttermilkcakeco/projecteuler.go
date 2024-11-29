package main

import (
	"fmt"
)

var singleDigits = [9]string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}
var tenToTwenty = [10]string{
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen",
	"eighteen", "nineteen",
}
var tens = [...]string{
	"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

func main() {

	oneToHundred := 0
	for _, word := range singleDigits {
		oneToHundred += len(word)
	}
	for _, word := range tenToTwenty {
		oneToHundred += len(word)
	}
	for _, tensWord := range tens {
		oneToHundred += len(tensWord)
		for _, onesWord := range singleDigits {
			oneToHundred += len(tensWord) + len(onesWord)
		}
	}

	total := oneToHundred
	for _, hundredsWord := range singleDigits {
		total += len(hundredsWord) + len("hundred")
		total += (len(hundredsWord)+len("hundredand"))*99 + oneToHundred
	}

	total += len("onethousand")

	fmt.Println(total)
}
