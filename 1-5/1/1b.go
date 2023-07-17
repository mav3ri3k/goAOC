package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkMax(max *[]int, calories int) {
  *max = append(*max, calories)
  sort.Ints(*max)
  *max = (*max)[1:]
}

func part2(max *[]int) {
  read,_ := os.ReadFile("1.txt")     
  inputString := string(read) 
  
  for _,elf := range strings.Split(inputString, "\n\n") {
    var calories int 
    for _,calorie := range strings.Split(elf, "\n") {
      value,_ := strconv.Atoi(calorie)
      calories += value
    }
    checkMax(max, calories)
  }
  totalCalories := 0
  for _,calorie := range *max {
    totalCalories += calorie
  }
  fmt.Println(totalCalories)
}

func main() {
  max := make([]int, 3)
  part2(&max)
}
