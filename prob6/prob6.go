package main

import (
	"fmt"
)

func main() {

	s1 := 0
	s2 := 0

	for i := 1; i <= 100; i++ {
		s1 += i * i
		s2 += i
	}

	s2 = s2 * s2
	fmt.Println(s2 - s1)
}
