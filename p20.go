package main

import (
  "fmt"
)

func fact (n int64) int64 {
  var result int64 = 1
  var step int64 = 1

  for ; step <= result; step++ {
    result = result * step
  }

  return result
}



// можно выкинуть все 10. тк они не влияют на сумму цифр

func main () {
  fmt.Println(fact(10))
}
