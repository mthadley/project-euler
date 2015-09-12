package main

import (
	"fmt"
	"math/big"
)

func main() {
	const gridSize = 20

	bigInt := big.NewInt(0)
	fmt.Println(bigInt.Binomial(2*gridSize, gridSize))
}
