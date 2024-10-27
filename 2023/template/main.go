package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var inputFile string

func main() {
  fmt.Printf("Part 01: %d\n", part1(inputFile))
  fmt.Printf("Part 02: %d\n", part2(inputFile))
}

func part1(input string) int64 {
  return 0
}

func part2(input string) int64 {
  return 0
}
