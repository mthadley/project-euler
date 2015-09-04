package main

import (
	"fmt"
	"math"
)

func main() {
	const number int = 600851475143

	largest := 0

	for i := 3; i <= sqrt(number); i += 2 {
		if number%i == 0 && isPrime(i) {
			largest = i
		}
	}

	fmt.Println(largest)
}

func isPrime(num int) bool {
	for i := 3; i <= sqrt(num); i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func sqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}
