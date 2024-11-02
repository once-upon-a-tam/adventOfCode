package adventOfCode_2023_05

import (
	"adventOfCode/helpers"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "05",
	Short: "Day 05 of Advent of Code 2023",
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

var (
	blankLineRegexp = regexp.MustCompile(`^$`)
)

type Range struct {
	start int
	end   int
}

func part1(input string) int64 {
	chunks := strings.Split(input, "\n\n")
	seedsLine, _ := strings.CutPrefix(chunks[0], "seeds:")

	seeds, err := helpers.IntsFromString(seedsLine, " ")
	if err != nil {
		panic(err)
	}

	var converted, todo []int
	converted = seeds
	for _, chunk := range chunks[1:] {
		conversionStrings := strings.Split(chunk, "\n")[1:]
		todo = append(todo, converted...)
		converted = nil

		for _, conversionStr := range conversionStrings {
			targetRange, delta, err := ParseConversionRule(conversionStr)
			if err != nil {
				panic(err)
			}
			var newConverted []int
			newConverted, todo = convert(todo, targetRange, delta)
			converted = append(converted, newConverted...)
		}
	}

	minLocation := math.MaxInt64
	all := append(converted, todo...)
	for _, location := range all {
		if location < minLocation {
			minLocation = location
		}
	}

	return int64(minLocation)
}

func part2(input string) int64 {
	chunks := strings.Split(input, "\n\n")

	// Parse the seeds line's values into ranges of seeds
	seedsLine, _ := strings.CutPrefix(chunks[0], "seeds:")
	seedValues, err := helpers.IntsFromString(seedsLine, " ")
	if err != nil {
		panic(err)
	}

	var seedsRanges []Range
	for i := 0; i < len(seedValues); i += 2 {
		seedsRanges = append(
			seedsRanges,
			Range{
				start: seedValues[i],
				end:   seedValues[i] + seedValues[i+1],
			},
		)
	}

	// We initialize the "converted" range slice with the seed range, then we
	// iterate over each section of the file (chunks[1:], as the first chunk is
	// the seeds range declaration)
	//
	// Each line of each section describes the conversion rule.
	// We iterate over the conversion rules, and convert the current range (todos)
	// using the convertRanges function, which either
	//    - applies the conversion to the part of the range that overlaps it,
	//    - appends the part that doesn't overlap to the "todos" slice
	//
	// We then end up with a list of location ranges
	var converted, todos []Range
	converted = seedsRanges
	for _, chunk := range chunks[1:] {
		conversionStrings := strings.Split(chunk, "\n")[1:]
		todos = append(todos, converted...)
		converted = nil

		for _, conversionStr := range conversionStrings {
			targetRange, delta, err := ParseConversionRule(conversionStr)
			if err != nil {
				panic(err)
			}
			var newConverted []Range
			newConverted, todos = convertRanges(todos, targetRange, delta)
			converted = append(converted, newConverted...)
		}
	}

	minLocation := math.MaxInt64
	all := append(converted, todos...)
	for _, location := range all {
		if location.start < minLocation {
			minLocation = location.start
		}
	}

	return int64(minLocation)
}

// ParseConversionRule parses a mapping rule into its target range and delta information.
func ParseConversionRule(s string) (targetRange Range, delta int, err error) {
	if blankLineRegexp.MatchString(s) {
		return Range{}, 0, nil
	}
	values, err := helpers.IntsFromString(s, " ")
	if err != nil {
		return Range{}, 0, err
	}

	if len(values) != 3 {
		return Range{}, 0, fmt.Errorf("invalid input: expected 3 numbers, got %d", len(values))
	}

	targetRange = Range{
		start: values[1],
		end:   values[1] + values[2] - 1,
	}

	delta = values[0] - values[1]

	return
}

// convert converts a slice of elements using a targetRange and a delta.
// If an element from the "elements" slice is within the "targetRange" range,
// we apply the "delta" value to it and append the result to "converted"
// If it doesn't, we append the initial value to "todos"
func convert(elements []int, targetRange Range, delta int) (converted []int, todo []int) {
	for _, element := range elements {
		if targetRange.start <= element && element <= targetRange.end {
			converted = append(converted, element+delta)
		} else {
			todo = append(todo, element)
		}
	}

	return
}

// convertRanges applies a conversion logic to the overlapping part of elements
// the range of conversionRule, and appends the rest of elements to the todos slice.
//
// If elements is fully overlapping with the conversionRule's range:
//    - converted is appended with the converted version of elements
//    - todos is untouched
//
// If elements doesn't overlap with the conversionRule's range:
//    - converted is untouched
//    - todos is appended with elements
//
// If elements partially overlaps with the conversionRule's range:
//    - converted is appended with the converted version of the overlapping part of elements
//    - todos is appended with the non-overlapping part of elements.

func convertRanges(elements []Range, targetRange Range, delta int) (converted []Range, todos []Range) {
	for _, element := range elements {
		if element.start > targetRange.end || element.end < targetRange.start {
			todos = append(todos, element)
			continue
		}

		// Fully overlapping
		if element.start >= targetRange.start && element.end <= targetRange.end {
			converted = append(
				converted,
				Range{
					start: element.start + delta,
					end:   element.end + delta,
				},
			)
			continue
		}

		// E     ------
		// C   -----
		// Partial overlapping with excess on the right
		if element.start <= targetRange.end && element.end > targetRange.end {
			todoRange := Range{
				start: targetRange.end + 1,
				end:   element.end,
			}

			convertedRange := Range{
				start: element.start + delta,
				end:   targetRange.end + delta,
			}
			// fmt.Printf("Partial overlapping, excess on right\t")
			// fmt.Printf("Converted: %d, todos: %d\n", Range{element.start,conversionRule.end}, todoRange)

			todos = append(todos, todoRange)
			converted = append(converted, convertedRange)

			continue
		}

		// E   -----
		// C     ------
		// Partial overlapping with excess on the left
		if element.start < targetRange.start && element.end >= targetRange.start {
			todoRange := Range{
				start: element.start,
				end:   targetRange.start - 1,
			}

			convertedRange := Range{
				start: targetRange.start + delta,
				end:   element.end + delta,
			}
			// fmt.Printf("Partial overlapping, excess on left\t")
			// fmt.Printf("Converted: %d, todos: %d\n", Range{conversionRule.start + conversionRule.delta, element.end + conversionRule.delta}, todoRange)

			todos = append(todos, todoRange)
			converted = append(converted, convertedRange)

			continue
		}
	}
	return
}
