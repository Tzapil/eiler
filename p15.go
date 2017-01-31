package main

import "fmt"


func main () {
    var table [20][20]int

    for i := 19; i >= 0; i-- {
        for j:= 19; j >= 0; j-- {
            right := 1
            down := 1

            if i < 19 {
                right = table[i + 1][j]
            }

            if j < 19 {
                down = table[i][j + 1]
            }

            table[i][j] = right + down
        }
    }

    fmt.Println(table[0][0])
}