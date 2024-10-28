package adventOfCode_2023_02

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
  onlyNumbers = regexp.MustCompile(`\d+`)
  onlyLetters = regexp.MustCompile(`[a-z]+`)
  blankString = regexp.MustCompile(`^$`)
)

//go:embed input.txt
var inputFile string

type Draw struct {
  Red int
  Green int
  Blue int
}

func main() {
  fmt.Printf("Part 01: %d\n", part1(inputFile))
  fmt.Printf("Part 02: %d\n", part2(inputFile))
}

func part1(input string) int64 {
  result := 0
  maximumDraw := Draw{12, 13, 14} // The total amount of each cube in the bag.

GameLoop:
  for _, gameStr := range strings.Split(input, "\n") {
    if blankString.MatchString(gameStr) {
      continue
    }
    // The first part is the game id, the second part is the list of draws
    gameParts := strings.Split(gameStr, ":")
    if len(gameParts) < 2 {
      log.Fatal("Invalid game format")
    }

    // Each draw combination is separated by a ";"
    for _, drawList := range strings.Split(gameParts[1], ";") {
      // Each amount-colour is separated by a ","
      for _, drawStr := range strings.Split(drawList, ",") {
        amount, err := strconv.Atoi(onlyNumbers.FindString(drawStr))
        if err != nil {
          log.Fatal(err)
        }
        color := onlyLetters.FindString(drawStr)

        switch {
        case color == "red" && amount > maximumDraw.Red:
          continue GameLoop
        case color == "green" && amount > maximumDraw.Green:
          continue GameLoop
        case color == "blue" && amount > maximumDraw.Blue: 
          continue GameLoop
        default:
        }
      }
    }

    gameId, err := strconv.Atoi(onlyNumbers.FindString(gameParts[0]))
    if err != nil {
      log.Fatal(err)
    }
    result += gameId
  }

  return int64(result)
}

func part2(input string) int64 {
  result := 0

  for _, gameStr := range strings.Split(input, "\n") {
    if blankString.MatchString(gameStr) {
      continue
    }
    // The first part is the game id, the second part is the list of draws
    gameParts := strings.Split(gameStr, ":")
    if len(gameParts) < 2 {
      log.Fatal("Invalid game format")
    }

    minimumAmount := Draw{}

    // Each draw combination is separated by a ";"
    for _, drawList := range strings.Split(gameParts[1], ";") {
      // Each amount-colour is separated by a ","
      for _, drawStr := range strings.Split(drawList, ",") {
        amount, err := strconv.Atoi(onlyNumbers.FindString(drawStr))
        if err != nil {
          log.Fatal(err)
        }
        color := onlyLetters.FindString(drawStr)

        switch {
        case color == "red" && amount > minimumAmount.Red:
          minimumAmount.Red = amount
        case color == "green" && amount > minimumAmount.Green:
          minimumAmount.Green= amount
        case color == "blue" && amount > minimumAmount.Blue: 
          minimumAmount.Blue= amount
        default:
        }
      }
    }

    result += minimumAmount.Red * minimumAmount.Green * minimumAmount.Blue
  }

  return int64(result)
}

