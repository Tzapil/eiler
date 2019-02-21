package main

import (
	"fmt"
	"math"
	"strconv"
)

func copy_slice(input []int64) []int64 {
	tmp := make([]int64, len(input))
	copy(tmp, input)
	return tmp
}

// func permutations(input []int64) [][]int64 {
// 	var result [][]int64

// 	if len(input) <= 1 {
// 		if len(input) == 1 {
// 			result = append(result, input)
// 		}
// 		return result
// 	}

// 	for i, element := range input {
// 		less_perm := permutations(append(copy_slice(input[0:i]), input[i+1:len(input)]...))
// 		for j, el := range less_perm {
// 			less_perm[j] = append([]int64{element}, el...)
// 		}

// 		result = append(result, less_perm...)
// 	}

// 	return result
// }

func rotations(input []int64) [][]int64 {
	var result [][]int64

	if len(input) >= 1 {
		result = append(result, input)
	}

	if len(input) <= 1 {

		return result
	}

	current := copy_slice(input)
	for i := 0; i < len(input); i++ {
		current = append(copy_slice(current[1:len(current)]), current[0])

		result = append(result, current)
	}

	return result
}

func isPrime(n int64, primes []int64) bool {
	sq := math.Sqrt(float64(n))
	for _, prime := range primes {
		if n == prime {
			return true
		}

		if math.Mod(float64(n), float64(prime)) == 0 {
			return false
		}

		// no use
		if float64(prime) > sq {
			break
		}
	}

	return true
}

func cut_numbers(n int64) []int64 {
	result := []int64{}

	s := strconv.Itoa(int(n))

	for _, char := range s {
		num, _ := strconv.Atoi(string(char))
		result = append(result, int64(num))
	}

	return result
}

func join_numbers(n []int64) int64 {
	var result int64 = 0
	var step int64 = 1
	length := len(n)

	for i := length - 1; i >= 0; i-- {
		result += step * n[i]
		step *= 10
	}

	return result
}

func main() {
	//1000000
	min := int64(1000000)

	fmt.Println(min)

	var result int64 = 0
	var i int64

	primes := []int64{2}

	for i = 3; i < min; i += 2 {
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}

	fmt.Println(primes)

	all := []int64{}

	for _, prime := range primes {
		nums := cut_numbers(prime)
		rotation := rotations(nums)

		var flag bool = true

		for _, p := range rotation {
			if !isPrime(join_numbers(p), primes) {
				flag = false
				break
			}
		}

		if flag {
			result++
			all = append(all, prime)
		}
	}

	fmt.Println(all)
	fmt.Println(result)
}
