package main

import (
	"fmt"
	"math"
	"strconv"
)

func permutations(input []int64) [][]int64 {
	var result [][]int64

	if len(input) <= 1 {
		result = append(result, input)
		return result
	}

	for _, element := range input {
		row := []int64{element}
		permutations
	}

	temp := strconv.Itoa(int(n))

	for _, char := range temp {
		num, _ := strconv.Atoi(string(char))
		result = append(result, int64(num))
	}

	return result
}

func isPrime(n int64, primes []int64) bool {
	for prime := range primes {
		if math.Mod(float64(n), float64(prime)) == 0 {
			return false
		}
	}

	return true
}

func main() {
	min := int64(1000000)

	fmt.Println(min)

	var result int64
	var primes []int64
	var i int64

	for i = 2; i < min; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)

			permutations := 
		}
	}

	fmt.Println(result)
}

// answer 40730
