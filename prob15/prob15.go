package main

import (
	"fmt"
	"math/big"
)

func Factorial(n int) *big.Int {
	f := big.NewInt(1)
	for i := 1; i <= n; i++ {
		f.Mul(f, big.NewInt(int64(i)))
	}
	return f
}

func main() {

	a := Factorial(40)
	b := Factorial(20)

	fmt.Println(a.Div(a, b.Mul(b, b)))
}
