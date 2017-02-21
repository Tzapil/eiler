package main

import "fmt"

func buf_x_2(b []int) {
  add := 0
  for i := 0; i < len(b); i++ {
    current := (b[i] * 2) + add
    b[i] = current % 10
    add = current / 10
  }
}

func main () {
  var buffer [350]int
  buffer[0] = 1

  for i := 0; i < 1000; i++ {
     buf_x_2(buffer[0:349])
  }

  result := 0
  for i := 0; i < len(buffer); i++ {
    result = result + buffer[i]
  }
  fmt.Println(result)
}
