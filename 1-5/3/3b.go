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
  input,_ := os.ReadFile("3.txt")
  totalPriorities := 0
  var rucksacks []string
    
  for _,rucksack := range strings.Split(string(input), "\n") {
    rucksacks = append(rucksacks, rucksack)
  }
  rucksacks = rucksacks[:len(rucksacks) - 1]
  for i1,rucksack := range rucksacks {
    if i1 % 3 != 0 {
      continue
    }
    for i2 := range rucksack {
      query1 := strings.Contains(rucksacks[i1 + 1], string(rucksack[i2]))
      query2 := strings.Contains(rucksacks[i1 + 2], string(rucksack[i2]))

      if query1 && query2{
        //fmt.Println(i1, string(rucksack[i2]), priority(rucksack[i2]))
        totalPriorities += priority(rucksack[i2])
        break
      }
    }
  }
  fmt.Println(totalPriorities)
}
