package main

import "fmt"

func main() {
	sum := 0

	for i := range fibGenerator() {
		if i < 4000000 {
			if i%2 == 0 {
				sum += i
			}
		} else {
			break
		}
	}

	fmt.Println(sum)
}

func fibGenerator() (c chan int) {
	c = make(chan int)

	x, y := 1, 2

	go func() {
		c <- x
		c <- y

		for {
			x, y = y, x+y

			c <- y
		}
	}()

	return c
}
