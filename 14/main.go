package main

import "fmt"

func main() {
	const limit = 1000000

	largest := 0
	largestCount := 0

	collatz := make([]int, limit)

	for i := 1; i < limit; i++ {
		count := 0

		for c := i; c > 1; {
			if c < limit && collatz[c] > 0 {
				count += collatz[c]
				break
			}

			count++

			if c%2 == 0 {
				c /= 2
			} else {
				c = 3*c + 1
			}
		}

		collatz[i] = count

		if count > largestCount {
			largest = i
			largestCount = count
		}
	}

	fmt.Println(largest)
}
