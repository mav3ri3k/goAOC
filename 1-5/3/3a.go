package main

import (
	"fmt"
	"os"
	"strings"
)

func priority(a byte) int {
  if a > 96 {
    return int(a - 'a' + 1)
  }
  return int(a - 'A' + 27)
}

func main() {
  rucksacks,_ := os.ReadFile("3.txt")
  totalPriorities := 0
  for _,rucksack := range strings.Split(string(rucksacks), "\n") {
    start := len(rucksack)/2
    for i := range rucksack[:start] {
      query := strings.ContainsAny(rucksack[start:], string(rucksack[i]))
      if i > start {
        break
      } else {
        if query {
          totalPriorities += priority(rucksack[i])
          break 
        }
      }
    }
  }
  fmt.Println(totalPriorities)
}
