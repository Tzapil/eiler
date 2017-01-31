package main

import "math"
import "fmt"


func div_count (n int64) int64 {
    var result int64 = 2
    var border int64 = int64(math.Sqrt(float64(n))) + 1
    var i int64

    for i = 2; i < border; i++ {
        if n % i == 0 {
            result = result + 2
        }
    }

    return result
}

func main() {
    var divisors int64 = 500
    var curr int64 = 1
    var i int64 = 2

    for {
        curr = curr + i
        r := div_count(curr)

        if r >= divisors {
            fmt.Println("FOUND:", curr, i, r)
            break
        }

        i++
    }
}