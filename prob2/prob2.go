package main

import "fmt"

func fib(max int) chan int {
    c := make(chan int)
    go func() {
        x, y := 0, 1
        for x < max {
            c <- x
            x, y = y, x+y
        }
        close(c)
    }()
    return c
}

func main() {
	sum := 0

    for n := range fib(4000000) {
		if n % 2 == 0 {
			sum += n
		}
	}

	fmt.Println(sum)
}