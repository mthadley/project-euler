package main

import (
	"fmt"
	"math"
)

func main() {
	const n = 2000000
	sum := 0

	for i := 0; i < n; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
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
