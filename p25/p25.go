package main

import (
	"fmt"
	"math"
)

func main() {
	min := int64(math.Pow10(999))

	fmt.Println(min)

	var f1 int64 = 1
	var f2 int64 = 2

	for f1 < min {
		swap := f1
		f1 = f2
		f2 += swap
	}

	fmt.Println(f1)
}
