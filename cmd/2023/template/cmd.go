package adventOfCode_2023_XX

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
  Use: "XX",
  Short: "Day XX of Advent of Code 2023",
  Run: func(cmd *cobra.Command, args []string) {
    execute(cmd.Parent().Name(), cmd.Name())
  },
}

func execute(parent string, command string) {
  b, err := os.ReadFile(fmt.Sprintf(`cmd/%s/%s/input.txt`, parent, command))

  if err != nil {
    panic(err)
  }
  
  fmt.Printf("Part 01: %d\n", part1(string(b)))
	fmt.Printf("Part 02: %d\n", part2(string(b)))
}

func part1(input string) int64 {
  return 0
}

func part2(input string) int64 {
  return 0
}
