package main

import (
	"fmt"
	"math"
)

const constant string = "constant"

func main() {
	fmt.Println(constant)
	//constant = "new constant" Cannot reAssign to constant
	const d = 3e20 * 10
	fmt.Println(math.Sin(d))
	fmt.Println(int64(math.Floor(d)))
}
