package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(2)
	n = n.Exp(n, big.NewInt(1000), nil)

	s := 0
	for _, c := range n.String() {
		s += int(c) - 48
	}

	fmt.Println(s)
}
