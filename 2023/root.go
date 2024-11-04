package adventOfCode_2023

import (
	adventOfCode_2023_01 "adventOfCode/2023/01"
	adventOfCode_2023_02 "adventOfCode/2023/02"
	adventOfCode_2023_03 "adventOfCode/2023/03"
	adventOfCode_2023_04 "adventOfCode/2023/04"
	adventOfCode_2023_05 "adventOfCode/2023/05"
	adventOfCode_2023_06 "adventOfCode/2023/06"
	adventOfCode_2023_07 "adventOfCode/2023/07"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
  Use: "2023",
  Short: "Solutions to the 2023 edition of Advent of Code",
  Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
  Cmd.AddCommand(adventOfCode_2023_01.Cmd)
  Cmd.AddCommand(adventOfCode_2023_02.Cmd)
  Cmd.AddCommand(adventOfCode_2023_03.Cmd)
  Cmd.AddCommand(adventOfCode_2023_04.Cmd)
  Cmd.AddCommand(adventOfCode_2023_05.Cmd)
  Cmd.AddCommand(adventOfCode_2023_06.Cmd)
  Cmd.AddCommand(adventOfCode_2023_07.Cmd)
}

func Execute() {
  if err := Cmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
