package main

import (
	"fmt"
	"math"
)

func main() {
	var result float64 = 0
	var max float64 = 10000000000

	for i := 1; i <= 1000; i++ {
		var temp float64 = 1
		for j := 0; j < i; j++ {
			temp = math.Floor(math.Mod(float64(temp*float64(i)), max))
		}
		result = math.Floor(math.Mod(float64(result+temp), max))
	}

	fmt.Println(int64(result))
}
