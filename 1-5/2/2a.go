package main

import (
	"fmt"
	"os"
	"strings"
)

func roundScore(a byte, b byte) int {

  points := map[byte]int{'X':1, 'Y':2, 'Z':3}
  opponent := map[byte]int{'A':1, 'B':2, 'C':3} 
  
  //draw
  if opponent[a] == points[b] {
    return points[b] + 3
  }

  switch a {
  case 'A'://rock
    if b == 'Y' {//paper
      return points[b] + 6
    } else {
      return points[b]
    }
  case 'B'://paper
    if b == 'Z' {//scissor
      return points[b] + 6
    } else {
      return points[b]
    }
  case 'C'://scissor
    if b == 'X' {//rock
      return points[b] + 6
    } else {
      return points[b]
    }
  }
  return 0
}
