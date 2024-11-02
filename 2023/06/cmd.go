package adventOfCode_2023_06

import (
	"adventOfCode/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "06",
	Short: "Day 06 of Advent of Code 2023",
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
	lines := strings.Split(input, "\n")
	timesLine, _ := strings.CutPrefix(lines[0], "Time:")
	times, err := helpers.IntsFromString(timesLine, " ")
	if err != nil {
		panic(err)
	}

	distancesLine, _ := strings.CutPrefix(lines[1], "Distance:")
	distances, err := helpers.IntsFromString(distancesLine, " ")
	if err != nil {
		panic(err)
	}

	var total int64 = 1

	for i := 0; i < len(times); i++ {
		var winningCases int64

		for velocity := 0; velocity < times[i]; velocity++ {
			// d = v * t
			remainingTime := times[i] - velocity
			if velocity*remainingTime > distances[i] {
				winningCases++
			}
		}
		total *= winningCases
	}
	return total
}

func part2(input string) int64 {
	lines := strings.Split(input, "\n")
	timesLine, _ := strings.CutPrefix(lines[0], "Time:")
	time, err := strconv.Atoi(strings.ReplaceAll(timesLine, " ", ""))
	if err != nil {
		panic(err)
	}

	distancesLine, _ := strings.CutPrefix(lines[1], "Distance:")
	distance, err := strconv.Atoi(strings.ReplaceAll(distancesLine, " ", ""))
	if err != nil {
		panic(err)
	}

	var winningCases int64
	for velocity := 0; velocity < time; velocity++ {
		remainingTime := time - velocity
		// d = v * t
		if velocity*remainingTime > distance {
			winningCases++
		}
	}

	return winningCases
}
