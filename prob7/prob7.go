package main

import (
	"fmt"

	"euler/utils"
)

func main() {

	n := 0
	for i := range utils.Primes(2, 1000000000) {
		n++
		if n == 10001 {
			fmt.Println(i)
			break
		}
	}
}
