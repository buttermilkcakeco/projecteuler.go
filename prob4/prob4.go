package main

import (
	"fmt"
)

func isPalindrome(n int) bool {
	s := fmt.Sprint(n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {

	largest := 0

	for a := 100; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			n := a * b
			if isPalindrome(n) {
				largest = max(n, largest)
			}
		}
	}

	fmt.Println(largest)
}
