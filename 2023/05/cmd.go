package adventOfCode_2023_05

import (
	"adventOfCode/helpers"
	"fmt"
	"os"
	"regexp"
	"sort"
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

func part1(input string) int64 {
	maps, err := almanacMapsFromInput(input)
	if err != nil {
		fmt.Println(err)
	}

  chunks := strings.Split(input, "\n\n")
	seedsRangesList, err := helpers.IntsFromString(strings.TrimPrefix(chunks[0], "seeds:"), " ")
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// To match with part 2, we use a list of intervals of range 1.
	seedRanges := make([]interval, 0)
	for _, seed := range seedsRangesList {
		seedRanges = append(seedRanges, interval{start: seed, end: seed})
	}

	sort.Slice(seedRanges, func(i, j int) bool {
		return seedRanges[i].start < seedRanges[j].start
	})

	minLocation := 0
	for _, seed := range seedRanges {
		location := findSeedRangeLocation(seed, maps)

		if minLocation == 0 {
			minLocation = location.start
		} else if minLocation > location.start {
			minLocation = location.start
		}
	}

	return int64(minLocation)
}

func part2(input string) int64 {
	maps, err := almanacMapsFromInput(input)
	if err != nil {
		fmt.Println(err)
	}

	chunks := strings.Split(input, "\n\n")
	seedsRangesList, err := helpers.IntsFromString(strings.TrimPrefix(chunks[0], "seeds:"), " ")
	if err != nil {
		fmt.Println(err)
		return 0
	}

	seedRanges := make([]interval, 0)
	for i := 0; i < len(seedsRangesList); i += 2 {
		seedRanges = append(
			seedRanges,
			interval{
				start: seedsRangesList[i],
				end:   seedsRangesList[i] + seedsRangesList[i+1],
			},
		)
	}

	sort.Slice(seedRanges, func(i, j int) bool {
		return seedRanges[i].start < seedRanges[j].start
	})

	for _, currMap := range maps {
		currMap.sortRanges()
		currMap.fillGapsInRanges()
	}

	minLocation := 0
	for _, seedRange := range seedRanges {
		location := findSeedRangeLocation(seedRange, maps)

		if minLocation == 0 {
			minLocation = location.start
		} else if minLocation > location.start {
			minLocation = location.start
		}
	}

	return int64(minLocation)
}

func findSeedRangeLocation(seeds interval, maps []almanacMap) interval {
	v := seeds

Maps:
	for _, mapping := range maps {
		for _, mapRange := range mapping.ranges {
			if v.start >= mapRange.match.start && v.end < mapRange.match.end {
				v.start = v.start + mapRange.offset
				v.end = v.end + mapRange.offset
				continue Maps
			}
		}
	}

	return v
}

type interval struct {
	start int // inclusive
	end   int // exclusive
}

var (
	categories = [...]string{
		"seed",
		"soil",
		"fertilizer",
		"water",
		"light",
		"temperature",
		"humidity",
		"location",
	}
)

type almanacMap struct {
	sourceCategory      string
	destinationCategory string
	ranges              []almanacRange
}

type almanacRange struct {
	match  interval
	offset int
}

func (m almanacMap) sortRanges() {
	sort.Slice(m.ranges, func(i, j int) bool {
		return m.ranges[i].match.start < m.ranges[j].match.start
	})
}

func (m *almanacMap) fillGapsInRanges() {
	updatedRanges := make([]almanacRange, 0)

	for _, currRange := range m.ranges {
		switch {
		case len(updatedRanges) == 0 && currRange.match.start != 0:
			updatedRanges = append(updatedRanges, almanacRange{
				match: interval{
					start: 0,
					end:   currRange.match.start,
				},
				offset: 0,
			})
		case len(updatedRanges) > 0 && currRange.match.start != updatedRanges[len(updatedRanges)-1].match.end:
			updatedRanges = append(updatedRanges, almanacRange{
				match: interval{
					start: updatedRanges[len(updatedRanges)-1].match.end,
					end:   currRange.match.start,
				},
				offset: 0,
			})
		}

		updatedRanges = append(updatedRanges, currRange)
	}

	m.ranges = updatedRanges
}

func almanacMapsFromInput(input string) ([]almanacMap, error) {
	chunks := strings.Split(input, "\n\n")

	if len(chunks) != len(categories) {
		return nil, fmt.Errorf("Expected %d categories, received %d", len(categories), len(chunks))
	}

	maps := make([]almanacMap, 0)
	for _, chunk := range chunks[1:] {
		parsed, err := almanacMapFromString(strings.Split(chunk, "\n"))
		if err != nil {
			return nil, err
		}
		maps = append(maps, parsed)
	}

	return maps, nil
}

func almanacMapFromString(s []string) (almanacMap, error) {
	if len(s) < 2 {
		return almanacMap{}, fmt.Errorf("invalid input: expected at least 2 lines, got %d", len(s))
	}

	source, destination, err := categoriesFromMapHeader(s[0])
	if err != nil {
		return almanacMap{}, fmt.Errorf("could not parse map header: %w", err)
	}

	var ranges []almanacRange
	for _, line := range s[1:] {
		if blankLineRegexp.MatchString(line) {
			continue
		}
		values, err := helpers.IntsFromString(line, " ")
		if err != nil {
			return almanacMap{}, err
		}

		if len(values) != 3 {
			return almanacMap{}, fmt.Errorf("invalid input: expected 3 numbers, got %d", len(values))
		}

		destStart := values[0]
		srcStart := values[1]
		length := values[2]

		ranges = append(
			ranges,
			almanacRange{
				match:  interval{start: srcStart, end: srcStart + length},
				offset: destStart - srcStart,
			})
	}

	return almanacMap{
		sourceCategory:      source,
		destinationCategory: destination,
		ranges:              ranges,
	}, nil
}

func categoriesFromMapHeader(s string) (string, string, error) {
	s = strings.TrimSuffix(s, " map:")
	parts := strings.Split(s, "-to-")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid input: expected 2 parts, got %d", len(parts))
	}
	return parts[0], parts[1], nil
}
