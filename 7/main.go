package main

import (
	"fmt"
	"math"
)

func main() {
	count, prime := 0, 0

	for i := 1; ; i++ {
		if isPrime(i) {
			prime = i
			count++
		}

		if count == 10001 {
			break
		}
	}

	fmt.Println(prime)
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
