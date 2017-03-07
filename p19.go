package main

import (
  "fmt"
)

func leap_year (year int) bool {
  return year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)
}

func next_month (year, month, day int) int {
  result := day
  if month == 2 {
    if (leap_year(year)) {
      result += 1
    }
  } else {
    result += 2
    if (month < 8 && month % 2 != 0) || (month >= 8 && month % 2 == 0){
      result += 1
    }
  }

  return result % 7
}

func main () {
  result := 0
  day := 1
  for year := 1901; year <= 2000; year++ {
    for month := 1; month <= 12; month++ {
      if (day == 6) {
        result += 1
      }
      day = next_month(year, month, day)
    }
  }

  fmt.Println(result)
}
