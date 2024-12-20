package adventOfCode_2023_01

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var (
	onlyNumbers    = regexp.MustCompile(`\d`)
	blankString    = regexp.MustCompile(`^$`)
)

var Cmd = &cobra.Command{
  Use: "01",
  Short: "Day 01 of Advent of Code 2023",
  Run: func(cmd *cobra.Command, args []string) {
    execute(cmd.Parent().Name(), cmd.Name())
  },
}

func execute(parent string, command string) {
  b, err := os.ReadFile(fmt.Sprintf(`%s/%s/input.txt`, parent, command))

  if err != nil {
    panic(err)
  }
  
  fmt.Printf("Part 01: %d\n", part1(string(b)))
	fmt.Printf("Part 02: %d\n", part2(string(b)))
}

func part1(input string) int64 {
	result := 0

	for _, entry := range strings.Split(input, "\n") {
		digits := onlyNumbers.FindAllString(entry, -1)

		switch len(digits) {
		case 0:
		case 1:
			number, _ := strconv.Atoi(digits[0])
			result += number * 11
		default:
			numberA, _ := strconv.Atoi(digits[0])
			numberB, _ := strconv.Atoi(digits[len(digits)-1])
			result += numberA*10 + numberB
		}
	}

	return int64(result)
}

func part2(input string) int64 {
	result := 0

  // WARNING: for some lines, such as "5tg578fldlcxponefourtwonet", the regexp
  // method doesn't work, as numbers overlap (twone contains both two and one")
	for _, entry := range strings.Split(input, "\n") {
		if blankString.MatchString(entry) {
			continue
		}
    firstDigit := findFirstDigit(entry)
		lastDigit := findLastDigit(entry)
    result += firstDigit * 10 + lastDigit
	}

	return int64(result)
}

var spelledDigits = map[string]int {
  "one": 1,
  "two": 2,
  "three": 3,
  "four": 4,
  "five": 5,
  "six": 6,
  "seven": 7,
  "eight": 8,
  "nine": 9,
}

func findFirstDigit(s string) int {
  for i:= 0; i < len(s); i++ {
    if found, d := containsSpelledDigit(s[:i]); found {
      return d
    } else if unicode.IsDigit(rune(s[i])) {
      return int(s[i] - '0')
    }
  }
  panic("No digit found in string" + s)
}

func findLastDigit(s string) int {
  for i:= len(s) - 1; i >= 0; i-- {
    if found, d := containsSpelledDigit(s[i:]); found {
      return d
    } else if unicode.IsDigit(rune(s[i])) {
      return int(s[i] - '0')
    }
  }
  panic("No digit found in string" + s)
}

func containsSpelledDigit(s string) (bool, int) {
	for k, v := range spelledDigits {
		if strings.Contains(s, k) {
			return true, v
		}
	}
	return false, 0
}

