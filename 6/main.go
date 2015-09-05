package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(squareSum(100) - sumSquares(100))
}

func squareSum(x int) (sum int) {
	sum = 0

	for i := 1; i <= x; i++ {
		sum += i
	}

	return pow(sum, 2)
}

func sumSquares(x int) (sum int) {
	sum = 0

	for i := 1; i <= x; i++ {
		sum += pow(i, 2)
	}

	return
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
