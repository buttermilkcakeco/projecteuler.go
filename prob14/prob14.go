package main

import (
	"fmt"
)

func main() {

	longestT, longestN := 0, 0

	for n := 1; n < 1000000; n++ {
		t, x := 1, n
		for x > 1 {
			if x%2 == 0 {
				x = x / 2
			} else {
				x = 3*x + 1
			}
			t++
		}

		if t > longestT {
			longestN, longestT = n, t
		}
	}

	fmt.Println(longestN, longestT)
}
