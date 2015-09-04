package main

import (
	"fmt"
	"strconv"
)

func main() {
	largest := 100000

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			k := i * j

			if k > largest && isPalindrome(k) {
				largest = k
			}
		}
	}

	fmt.Println(largest)
}

func isPalindrome(x int) bool {
	s := ""

	for i := x; i > 0; i /= 10 {
		s += strconv.Itoa(i % 10)
	}

	return strconv.Itoa(x) == s
}
