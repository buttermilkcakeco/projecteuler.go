package main

import (
	"fmt"
)

func main() {

	for a := 1; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			for c := b; c < 1000; c++ {
				if a*a+b*b == c*c && a+b+c == 1000 {
					fmt.Println(a, b, c, a*b*c)
					break
				}
			}
		}
	}
}
