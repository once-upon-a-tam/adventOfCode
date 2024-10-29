package main

import (
	adventOfCode_template "adventOfCode/template"
	adventOfCode_2023 "adventOfCode/2023"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use: "Advent of code",
  Short: "Solutions to various editions of the Advent of Code event",
  Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
  rootCmd.AddCommand(adventOfCode_template.Cmd)
  rootCmd.AddCommand(adventOfCode_2023.Cmd)
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func main() {
  rootCmd.Execute()
}

