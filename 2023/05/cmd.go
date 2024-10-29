package adventOfCode_2023_05

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
  Use: "05",
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
  seedListStr := GetStringInBetween(input, "seeds:", "seed-to-soil map:")
  seeds, err := parseMappingData(seedListStr)
  if err != nil {
    panic(err)
  }

  seedToSoilStr := GetStringInBetween(input, "seed-to-soil map:", "soil-to-fertilizer map:")
  seedToSoilMapping, err := parseMappingData(seedToSoilStr)
  if err != nil {
    panic(err)
  }

  soilToFertilizerStr := GetStringInBetween(input, "soil-to-fertilizer map:", "fertilizer-to-water map:")
  soilToFertilizerMapping, err := parseMappingData(soilToFertilizerStr)
  if err != nil {
    panic(err)
  }

  fertilizertoWaterStr := GetStringInBetween(input, "fertilizer-to-water map:", "water-to-light map:")
  fertilizerToWaterMapping, err := parseMappingData(fertilizertoWaterStr)
  if err != nil {
    panic(err)
  }
  
  waterToLightStr := GetStringInBetween(input, "water-to-light map:", "light-to-temperature map:")
  waterToLightMapping, err := parseMappingData(waterToLightStr)
  if err != nil {
    panic(err)
  }

  lightToTempStr:= GetStringInBetween(input, "light-to-temperature map:", "temperature-to-humidity map:")
  lightToTempMapping, err := parseMappingData(lightToTempStr)
  if err != nil {
    panic(err)
  }

  tempToHumidityStr := GetStringInBetween(input, "temperature-to-humidity map:", "humidity-to-location map:")
  tempToHumidityMapping, err := parseMappingData(tempToHumidityStr)
  if err != nil {
    panic(err)
  }

  _, humidityToLocationStr, _ := strings.Cut(input, "humidity-to-location map:")
  humidityToLocationMapping, err := parseMappingData(humidityToLocationStr)
  if err != nil {
    panic(err)
  }

  locations := make([]int, 0)
  for _, seed := range seeds[0] {
    soil, err := findMappingDestination(seed, seedToSoilMapping)
    if err != nil {
      panic(err)
    }

    fertilizer, err := findMappingDestination(soil, soilToFertilizerMapping)
    if err != nil {
      panic(err)
    }

    water, err := findMappingDestination(fertilizer, fertilizerToWaterMapping)
    if err != nil {
      panic(err)
    }

    light, err := findMappingDestination(water, waterToLightMapping)
    if err != nil {
      panic(err)
    }

    temp, err := findMappingDestination(light, lightToTempMapping)
    if err != nil {
      panic(err)
    }

    humidity, err := findMappingDestination(temp, tempToHumidityMapping)
    if err != nil {
      panic(err)
    }
    
    location, err := findMappingDestination(humidity, humidityToLocationMapping)
    if err != nil {
      panic(err)
    } 

    locations = append(locations, location)
  }

  lowestLocation := locations[0]
  for _, v := range locations {
    lowestLocation = min(lowestLocation, v)
  }

  return int64(lowestLocation)
}

// findMappingDestination returns the mapped destination related to the provided
// source inside the provided mappingData.
// If no destination is found, we return the provided source.
func findMappingDestination(source int, mappingData [][]int) (int, error) {
  for i := 0; i < len(mappingData); i++ {
    mapSource := mappingData[i][1]
    mapRange := mappingData[i][2]

    // If the source is in range of the current map, then we have a correspondence
    if source >= mapSource && source <= mapSource + mapRange {
      delta := mappingData[i][0] - mapSource
      return source + delta, nil
    }
  }

  return source, nil
}

func parseMappingData(input string) ([][]int, error) {
  result := make([][]int, 0)
  for _, str := range strings.Split(input, "\n") {
    if blankLineRegexp.MatchString(str) {
      continue
    }

    fields := strings.Fields(str)
    ints := make([]int, len(fields))

    for i, s := range fields {
      v, err := strconv.Atoi(s)
      if err != nil {
        return nil, err
      }
      ints[i] = v
    }
    result = append(result, ints)
  }

  return result, nil
}

// GetStringInBetween returns empty string if no start or end string found
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	return str[s : s+e]
}


func part2(input string) int64 {
  return 0
}
