package adventOfCode_2023_08

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "08",
	Short: "Day 08 of Advent of Code 2023",
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
	sections := strings.Split(input, "\n\n")
	instructions := strings.Split(sections[0], "")

	nodes := make(map[string][2]string)
	for _, line := range strings.Split(sections[1], "\n") {
		if regexp.MustCompile(`^$`).MatchString(line) {
			continue
		}
		start, left, right := ParseNode(line)
		nodes[start] = [2]string{left, right}
	}

	return FindStepsToReach("AAA", instructions, nodes)
}

// For each starting step, there's only one **Z node before it loops, and the
// loop destination is always N steps from the start, while the point it loops
// from is always N steps after the **Z node.
// We cane therefore get the minimum amount of steps to reach the desired state
// by calculating the Least Common Multiplier of each path.
func part2(input string) int64 {
	sections := strings.Split(input, "\n\n")
	instructions := strings.Split(sections[0], "")

	initialSteps := make([]string, 0)
	nodes := make(map[string][2]string)

	// Process the file to create a map of nodes
	for _, line := range strings.Split(sections[1], "\n") {
		if regexp.MustCompile(`^$`).MatchString(line) {
			continue
		}
		start, left, right := ParseNode(line)
		nodes[start] = [2]string{left, right}

		if start[len(start)-1] == 'A' {
			initialSteps = append(initialSteps, start)
		}
	}

	totalSteps := make([]int64, 0)

	// Calculate the minimum steps for each path from the starting steps to their
	// respective desired state.
	for _, step := range initialSteps {
		totalSteps = append(totalSteps, FindStepsToReach(step, instructions, nodes))
	}

	return LCM(totalSteps[0], totalSteps[1], totalSteps...)
}

// FindStepsToReach returns the amount of steps required to reach a step ending
// with "Z" from the provided startingStep.
func FindStepsToReach(startingStep string, instructions []string, nodes map[string][2]string) int64 {
	var totalSteps int64 = 0
	currentStep := startingStep
	nextInstruction := 0

Walk:
	for {
    // Was 'currentStep == 'ZZZ' in part1, got updated to match with part2.
		if currentStep[len(currentStep)-1] == 'Z' {
			break Walk
		}
		if instructions[nextInstruction] == string('R') {
			currentStep = nodes[currentStep][1]
		} else {
			currentStep = nodes[currentStep][0]
		}
		totalSteps++
		nextInstruction = (nextInstruction + 1) % len(instructions)
	}

	return totalSteps
}

func ParseNode(node string) (start, left, right string) {
	parts := strings.Split(node, " = ")
	dest := strings.Split(strings.Trim(parts[1], `()`), ", ")

	start = parts[0]
	left = dest[0]
	right = dest[1]

	return start, left, right
}

// GCD returns the greatest common divisor (GCD) of a and b via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM returns the Least Common Multiple (LCM) via GCD of a list of int64
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
