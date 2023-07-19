package main

import (
  "os"
  "fmt"
  "strings"
)

func result(a string, b string) int {
  results := map[string]int{
    "AY": 4,
    "AX": 3,
    "AZ": 8,
    "BY": 5,
    "BX": 1,
    "BZ": 9,
    "CY": 6,
    "CX": 2,
    "CZ": 7,
  }

  return results[a + b]
}

func main() {
  rounds,_ := os.ReadFile("2.txt")
  var totalScore int 
  for _,round := range strings.Split(string(rounds), "\n") {
    score := 0 
    if round != "" {
      score = result(round[:1], round[2:3]) 
    } 
    totalScore += score
  }
  fmt.Println(totalScore)
}
