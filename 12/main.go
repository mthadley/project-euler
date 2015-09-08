package main

import (
	"fmt"
	"math"
)

func main() {
root:
	for f := range tGenerator() {
		count := 2

		for i := 2; i <= int(math.Sqrt(float64(f))); i++ {
			if f%i == 0 {
				count += 2
			}

			if count > 500 {
				fmt.Println(f)
				break root
			}
		}
	}
}

func tGenerator() chan int {
	c := make(chan int)

	go func() {
		f := 1

		for i := 2; ; i++ {
			c <- f
			f += i
		}
	}()

	return c
}
