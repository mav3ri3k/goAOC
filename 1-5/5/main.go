package main

import (
	"fmt"
	"os"
	"strings"
)

func part1() {
	input, _ := os.ReadFile("input.txt")
  for _,stack := range strings.Split(string(input), "\n") {
    if stack == "" {
      break
    } 
    fmt.Println(stack)
  }
}

func main() {
	part1()
}
