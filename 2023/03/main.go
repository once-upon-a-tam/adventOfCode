package adventOfCode_2023_03

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

var (
  onlySymbols = regexp.MustCompile(`[^.0-9]`)
  onlyAsterisks = regexp.MustCompile(`\*`)
  onlyNumbers = regexp.MustCompile(`\d+`)
)

type Coordinates struct {
  row int
  col int
}

func main() {
  fmt.Printf("Part 01: %d\n", part1(inputFile))
  fmt.Printf("Part 02: %d\n", part2(inputFile))
}

func part1(input string) int64 {
  result := 0
  
  symbolsLocations := make([]Coordinates, 0)
  for i, row := range strings.Split(input, "\n"){
    // Find the x and y coordinates of each symbol, excluding dots
    for _, match := range onlySymbols.FindAllStringIndex(row, -1) {
      symbolsLocations = append(symbolsLocations, Coordinates{row: i, col: match[0]})
    }
  }

  // For each symbol, check if there is any number adjacent to it.
  // Add the sum of all the numbers adjacent to a symbol.
  for _, coordinates := range symbolsLocations {
    for _, number := range GetNumbersAdjacentToCoordinates(coordinates, strings.Split(input, "\n")) {
      result += number
    }
  }

  return int64(result)
}

func part2(input string) int64 {
  result := 0

  starsLocations := make([]Coordinates, 0)
  for i, row := range strings.Split(input, "\n") {
    // Find the x and y coordinates of each asterisk character.
    for _, match := range onlyAsterisks.FindAllStringIndex(row, -1) {
      starsLocations = append(starsLocations, Coordinates{row: i, col: match[0]})
    }
  }

  // For each asterisk, check if there are exactly 2 numbers adjacent to it.
  // If so, add their product to the total value.
  for _, coordinates := range starsLocations {
    adjacent := GetNumbersAdjacentToCoordinates(coordinates, strings.Split(input, "\n"))
    if len(adjacent) == 2 {
      result += adjacent[0] * adjacent[1]
    }
  }

  return int64(result)
}

// GetNumbersAdjacentToCoordinates returns a list of numbers that are adjacent
// to the provided coordinates, including diagonally.
func GetNumbersAdjacentToCoordinates(coord Coordinates, s []string) []int {
	numbers := make([]int, 0)
	searchRange := struct {
		row_start int
		row_end   int
		col_start int
		col_end   int
	}{
		max(coord.row-1, 0),
		min(coord.row+1, len(s)-1),
		coord.col,
		min(coord.col+1, len(s[0])),
	}

	for i := searchRange.row_start; i <= searchRange.row_end; i++ {
		// Find all numbers in the current row
		numbersInRow := onlyNumbers.FindAllStringIndex(s[i], -1)
		for _, match := range numbersInRow {
			// Check if the number's first index and last index range overlaps with
			// the search range.
			match_start := match[0]
			match_end := match[1]
			if searchRange.col_start <= match_end && match_start <= searchRange.col_end {
				number, err := strconv.Atoi(s[i][match_start:match_end])
				if err != nil {
					fmt.Println(err)
					continue
				}

				numbers = append(numbers, number)
			}
		}
	}

	return numbers
}


