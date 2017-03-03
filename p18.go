package main

import (
  "fmt"
  "bufio"
  "io"
  //"strconv"
  "os"
)

var pyramid [][]int = [][]int{
  []int{75,},
  []int{95, 64,},
  []int{17, 47, 82,},
  []int{18, 35, 87, 10,},
  []int{20, 04, 82, 47, 65,},
  []int{19, 01, 23, 75, 03, 34,},
  []int{88, 02, 77, 73, 07, 63, 67,},
  []int{99, 65, 04, 28, 06, 16, 70, 92,},
  []int{41, 41, 26, 56, 83, 40, 80, 70, 33,},
  []int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29,},
  []int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14,},
  []int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57,},
  []int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48,},
  []int{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31,},
  []int{04, 62, 98, 27, 23,  9, 70, 98, 73, 93, 38, 53, 60, 04, 23,},
}

func max2(a, b int) int {
  result := a
  if b > a {
    result = b
  }

  return result
}

func read_line(r io.Reader, count int) ([]int, error) {
  var result []int = make([]int, count)
  var e error = nil
  for i := 0; i < count; i++ {
    _, err := fmt.Fscan(r, &result[i])

    if err != nil {
      e = err
      break
    }
  }

  return result, e
}

func read_data(filename string) [][]int {
  var result [][]int

  // Open file
  file, oerr := os.Open(filename)
  if oerr != nil {
    panic(fmt.Errorf("The file %s does not exist!", filename))
  }

  // Close file at the end
  defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

  reader := bufio.NewReader(file)
  counter := 1
  for {
      arr, err := read_line(reader, counter)

      if err != nil {
        break
      }

      result = append(result, arr)

      counter += 1
  }

  return result
}

func main () {
  pyramid := read_data("./67data.txt")
  for i := len(pyramid) - 2; i >= 0; i-- {
    for j := 0; j < i + 1; j++ {
      pyramid[i][j] += max2(pyramid[i + 1][j], pyramid[i + 1][j + 1])
    }
  }

  fmt.Println(pyramid[0][0])
}
