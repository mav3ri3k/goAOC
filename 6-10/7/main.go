package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println(string(input))
}
