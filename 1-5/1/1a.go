package main

import (
	"fmt"
	"os"
	"strings"
  "strconv"
	//  "strings"
)

var elfs map[int]int = make(map[int]int)

func part1() {
  read,_ := os.ReadFile("1.txt")     
  inputString := string(read) 
  
  var max int = 0

  for _,elf := range strings.Split(inputString, "\n\n") {
    var calories int 
    for _,calorie := range strings.Split(elf, "\n") {
      value,_ := strconv.Atoi(calorie)
      calories += value
    }
    if max < calories {
      max = calories
    } 
  }
  fmt.Println(max)
}
