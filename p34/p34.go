package main

import (
	"fmt"
	"math"
	"strconv"
)

func digits(n int64) []int64 {
	var result []int64

	temp := strconv.Itoa(int(n))

	for _, char := range temp {
		num, _ := strconv.Atoi(string(char))
		result = append(result, int64(num))
	}

	return result
}

func fact(n int64) int64 {
	var result int64 = 1
	var step int64 = 1

	for ; step <= n; step++ {
		result = result * step
	}

	return result
}

func main() {
	min := int64(math.Pow10(8))

	fmt.Println(min)

	var result int64
	var facts [10]int64
	var i int64

	for i = 0; i < 10; i++ {
		facts[i] = fact(i)
	}

	fmt.Println(facts)

	for i = 10; i < min; i++ {
		var sum int64 = 0
		for _, d := range digits(i) {
			// fmt.Println(d, facts[d])
			sum += facts[d]
		}

		// fmt.Println(sum, result)
		if sum == i {
			result += sum
		}
	}

	fmt.Println(result)
}

// answer 40730
