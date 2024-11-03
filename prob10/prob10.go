package main

import (
	"euler/utils"
	"fmt"
)

func main() {

	s := 0
	for p := range utils.Primes(2, 2000000) {
		s += p
	}

	fmt.Println(s)
}
