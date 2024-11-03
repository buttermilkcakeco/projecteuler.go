package main

import (
    "fmt"

    "euler/utils"
)

func main() {
	largest := 0

    for n := range utils.Factorization(600851475143) {
		largest = n
	}

	fmt.Println(largest)
}