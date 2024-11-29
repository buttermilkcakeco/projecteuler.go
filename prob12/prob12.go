package main

import (
	"euler/utils"
	"fmt"
)

func TriangleNumbers() chan int {
	c := make(chan int)

	go func() {
		t := 1
		for i := 2; true; i++ {
			c <- t
			t += i
		}
		close(c)
	}()

	return c
}

func main() {

	for n := range TriangleNumbers() {

		s := 2
		for range utils.Factors(n) {
			s++
		}

		if s > 500 {
			fmt.Println(n, s)
			break
		}
	}

}
