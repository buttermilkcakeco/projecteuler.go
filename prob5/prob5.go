package main

import (
	"fmt"
)

func match(n int) bool {
	for i := 2; i < 20; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func main() {

	for n := 20; true; n++ {
		if match(n) {
			fmt.Println(n)
			break
		}
	}
}
