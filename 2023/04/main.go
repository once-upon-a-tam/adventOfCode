package adventOfCode_2023_04

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	blankString = regexp.MustCompile(`^$`)
)

//go:embed input.txt
var inputFile string

func main() {
	fmt.Printf("Part 01: %d\n", part1(inputFile))
	fmt.Printf("Part 02: %d\n", part2(inputFile))
}

func part1(input string) int64 {
	var score int64 = 0

	for _, cardStr := range strings.Split(input, "\n") {
		if blankString.MatchString(cardStr) {
			continue
		}

		matches := 0
		winning, drawn, err := ProcessCardString(cardStr)
		if err != nil {
			panic(err)
		}

		for _, w := range winning {
			if slices.Contains(drawn, w) {
				matches += 1
				continue
			}
		}

		score += int64(math.Pow(2, float64(matches-1)))
	}
	return score
}

func part2(input string) int64 {
	var score int64 = 0

  // matches stores the amount of matching number for each card.
  // The index of the card is equivalent to its Id - 1
  // eg: Card 1's matching numbers is stored in matches[0]
  var matches []int

  // copies stores the amount of copies for each card.
  // The index of the card is equivalent to its Id - 1
  // eg: Card 1's amount of copies is stored in copies[0]
  var copies []int
  

  matches = make([]int, 0)
	for _, cardStr := range strings.Split(input, "\n") {
		if blankString.MatchString(cardStr) {
			continue
		}

		matchingNumbers := 0
		winning, drawn, err := ProcessCardString(cardStr)
		if err != nil {
			panic(err)
		}

		for _, w := range winning {
			if slices.Contains(drawn, w) {
				matchingNumbers += 1
				continue
			}
		}
    matches = append(matches, matchingNumbers)
	}

  copies = make([]int, len(matches))

  // Each card has at least one copy (the original)
  for i := range matches {
    copies[i] = 1
  }

  // We iterate over each card.
  for cardIndex, matchingNumbers := range matches {
    // if Card 3 has 4 winning numbers, and we have 2 copies of it,
    // then we add a copy of Card 4, 5, 6, 7 twice
    for i := 1; i <= matchingNumbers; i ++ {
      copies[cardIndex + i] += copies[cardIndex] 
    }
  }
  
  for _, v := range copies {
    score += int64(v)
  }

  return score
}

// ProcessCardString splits the provided string into usable values
// Given the following string :
//
//	Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
//
// It will return a []int value for the winning numbers
// and another []int value for the list of numbers we have.
func ProcessCardString(s string) (Winning []int, Drawn []int, err error) {
	firstSplit := strings.Split(s, ":")
	numbers := strings.Split(firstSplit[1], "|")

	for _, winStr := range strings.Fields(numbers[0]) {
		winNbr, err := strconv.Atoi(winStr)
		if err != nil {
			return nil, nil, err
		}
		Winning = append(Winning, winNbr)
	}

	for _, drawStr := range strings.Fields(numbers[1]) {
		drawNbr, err := strconv.Atoi(drawStr)
		if err != nil {
			return nil, nil, err
		}
		Drawn = append(Drawn, drawNbr)
	}

	return Winning, Drawn, nil
}
