package utils

import (
	"math"
)

func Factors(n int) chan int {
	c := make(chan int)

	go func() {
		m := int(math.Sqrt(float64(n)))
		for i := 2; i <= m; i++ {
			if n%i == 0 {
				c <- i
				c <- n / i
			}
		}
		close(c)
	}()

	return c
}

func IsPrime(n int) bool {
	m := int(math.Sqrt(float64(n)))
	for i := 2; i <= m; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Primes(start int, stop int) chan int {
	c := make(chan int)

	go func() {
		for i := start; i < stop; i++ {
			if IsPrime(i) {
				c <- i
			}
		}
		close(c)
	}()

	return c
}

func Factorization(n int) chan int {
	c := make(chan int)

	go func() {
		m := int(math.Sqrt(float64(n)))
		for i := 2; i <= m; i++ {
			if n%i == 0 && IsPrime(i) {
				c <- i
			}
		}
		close(c)
	}()

	return c
}
