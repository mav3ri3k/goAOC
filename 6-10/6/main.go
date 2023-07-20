package main

import (
	"fmt"
	"os"
	"strings"
)

func part1() {
	input, _ := os.ReadFile("input.txt")
  str := string(input)
  for i := range str {
    repeat := 0
    for i1 := 0; i1 < 4; i1++ {
      condition := strings.Contains(str[i:i + i1] + str[i + i1 + 1:i+4], string(str[i + i1]))     
      if condition {
        repeat++
      }
    }
    if repeat == 0 {
      fmt.Println(i + 4) 
      break
    }
  }
}

func part2() {
	input, _ := os.ReadFile("input.txt")
  str := string(input)
  for i := range str {
    repeat := 0
    for i1 := 0; i1 < 14; i1++ {
      condition := strings.Contains(str[i:i + i1] + str[i + i1 + 1:i+14], string(str[i + i1]))     
      if condition {
        repeat++
      }
    }
    if repeat == 0 {
      fmt.Println(i + 14) 
      break
    }
  }
}

func main() {
  part2()
}
