package main

import "fmt"

func main() {
	const n = 1000

root:
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			k := n - i - j

			if i*i+j*j == k*k && i+j+k == n {
				fmt.Println(i * j * k)

				break root
			}
		}
	}
}
