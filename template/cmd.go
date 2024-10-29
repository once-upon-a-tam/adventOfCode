package adventOfCode_template

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
  Use: "new",
  Short: "Create a new directory for an advent of code daily challenge",
  Run: func(cmd *cobra.Command, args []string) {
    day, year, err := parseArgs(args)
    if err != nil {
      fmt.Println(err)
      return
    }

    path, err := setupNewDay(year, day);
    if err != nil {
      fmt.Println(err)
      return
    }

    fmt.Printf("New directory added at %s\n", path)
  },
}

func parseArgs(args []string) (string, string, error) {
  if len(args) < 2 {
    return "", "", errors.New("Usage: new YYYY DD")
  }

  year, err := strconv.Atoi(args[0])
  if err != nil {
    return "", "", err
  }

  day, err := strconv.Atoi(args[1])
  if err != nil {
    return "", "", err
  }

  dayStr := fmt.Sprintf("%02d", day)
  yearStr := fmt.Sprintf("%d", year)

  return dayStr, yearStr, nil
}

func setupNewDay(year string, day string) (dirPath string, err error) {
  dirPath = fmt.Sprintf(`%s/%s`, year, day)

  if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
    return "", fmt.Errorf(`the directory already exists at %s`, dirPath)
  }

  defer func() {
    // If an error occurred, remove the files.
    if err != nil && os.IsExist(err) {
      os.RemoveAll(dirPath)
    }
  }()

  vars := make(map[string]interface{})
  vars["Year"] = year
  vars["Day"] = day
  
  if err := os.Mkdir(dirPath, 0755); err != nil {
    return "", err
  }

  if err := createFromTemplate(
    "template/cmd.tmpl",
    fmt.Sprintf("%s/%s", dirPath, "cmd.go"),
    vars,
  ); err != nil {
    return "", err
  }

  if err := createFromTemplate(
    "template/cmd_test.tmpl",
    fmt.Sprintf("%s/%s", dirPath, "cmd_test.go"),
    vars,
  ); err != nil {
    return "", err
  }

  if _, err := os.Create(fmt.Sprintf(`%s/%s`, dirPath, "input.txt")); err != nil {
    return "", err
  }

  if _, err := os.Create(fmt.Sprintf(`%s/%s`, dirPath, "input_test.txt")); err != nil {
    return "", err
  }

  return dirPath, nil
}

func createFromTemplate(tmpl string, path string, vars map[string]interface{}) error {
  tmp, err := template.ParseFiles(tmpl)
  if err != nil {
    return err
  }
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  if err := tmp.Execute(file, vars); err != nil {
    return err
  }

  return nil
}
