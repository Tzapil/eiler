package main

import "fmt"

func next(n int64) int64 {
    var result int64 = n

    if n % 2 == 0 {
        result = n / 2
    } else {
        result = n * 3 + 1
    }

    return result
}

func chain_l(n int64) int64 {
    var result int64 = 1
    for n > 1 {
        result++
        n = next(n)
    }

    return result
}

func main () {
    var border int64 = 1000000
    var max int64 = 0
    var result int64 = 0

    var i int64

    for i = 1; i < border; i++ {
        l := chain_l(i)
        if l > max {
            max = l
            result = i
        }
    }

    fmt.Println(result, max)
}