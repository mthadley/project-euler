package main

import (
	"fmt"
	"math/big"
	//"strconv"
)

func main() {
	x := big.NewInt(2)
	y := big.NewInt(1000)

	exp := x.Exp(x, y, nil).String()

	sum := 0

	for _, r := range exp {
		sum += int(r - '0')
	}

	fmt.Println(sum)
}