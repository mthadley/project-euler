package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	file, _ := os.Open("numbers.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := big.NewInt(0)

	for scanner.Scan() {
		myInt := big.NewInt(0)
		myInt.SetString(scanner.Text(), 10)

		sum.Add(sum, myInt)
	}

	fmt.Println(sum)
}
