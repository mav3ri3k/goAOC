package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	input, _ := os.ReadFile("4.txt")
	counter := 0
	for _, pairs := range strings.Split(string(input), "\n") {
		if pairs == "" {
			continue
		}
		var values []int
		for _, pair := range strings.Split(pairs, ",") {
			str1 := strings.Split(pair, "-")[0]
			num1, _ := strconv.Atoi(str1)
			str2 := strings.Split(pair, "-")[1]
			num2, _ := strconv.Atoi(str2)
			values = append(values, num1, num2)
		}
		if values[0] <= values[2] && values[1] >= values[3] {
			counter += 1
		} else if values[0] >= values[2] && values[1] <= values[3] {
			counter += 1
		}
  
	}
	fmt.Println(counter)
}

func part2() {
	input, _ := os.ReadFile("4.txt")
	counter := 0
	for _, pairs := range strings.Split(string(input), "\n") {
		if pairs == "" {
			continue
		}
		var values []int
		for _, pair := range strings.Split(pairs, ",") {
			str1 := strings.Split(pair, "-")[0]
			num1, _ := strconv.Atoi(str1)
			str2 := strings.Split(pair, "-")[1]
			num2, _ := strconv.Atoi(str2)
			values = append(values, num1, num2)
		}
    case1 := values[2] >= values[0] && values[2] <= values[1]
    case2 := values[0] >= values[2] && values[0] <= values[3]
		if case1 || case2 {
      counter += 1
		} else {
    }
	}
  fmt.Println(counter)
}

func main() {
  part1() 
  part2()
}
