package main

import "fmt"

func main() {
	for i := 1; ; i++ {
		b := true

		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				b = false
				break
			}
		}

		if b {
			fmt.Println(i)
			break
		}
	}
}
