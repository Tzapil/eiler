package main

import "fmt"

// one two three four five six seven eight nine
// ten eleven twelve thirteen fourteen fifteen sixteen seventeen eighteen nineteen
// twenty thirty forty fifty sixty seventy eighty ninety
// hundred thousand

// form 1 to 999
func calc_len(n int) int {
  u := [10]int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4} // 0 to 9
  d := [10]int{0, 0, 6, 6, 5, 5, 5, 7, 6, 6} // 20 to 90
  d1 := [10]int{3, 6, 6, 8, 8, 7, 7, 9, 8, 8} // 10 to 19
  result := 0

  step := n
  state := 1
  for step != 0 {
    digit := step % 10

    switch state {
      case 1:
        result = u[digit]
        state = 2
      case 2:
        if digit == 1 {
          result = d1[n % 10]
        } else {
          result += d[digit]
        }

        state = 3
      case 3:
        if result != 0 {
          result = result + 3 // and length
        }
        result += + u[digit] + 7 // hundred length
    }

    step = step / 10
  }

  return result
}

func main () {
  result := 0
  for i := 1; i < 1000; i++ {
    fmt.Println(i, calc_len(i))
    result += calc_len(i)
  }
  fmt.Println(result + 11)
}
