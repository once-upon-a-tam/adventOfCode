# 🎄 Advent of Code 🎄

Working through problems for the various [Advent of Code](https://adventofcode.com)
events.

Each year has its own dedicated subfolder within this project :

- [2023](2023), in [Go](https://go.dev/)

## Tools

- [Go](https://go.dev/) as the programming language
- [Cobra](https://github.com/spf13/cobra) to provide the command-line interface
- [Testify](https://github.com/stretchr/testify) to help with testing

## How to run the solutions

This repository makes use of [Cobra](https://github.com/spf13/cobra) to provide
the command-line interface.
To run the solution for a puzzle, use `go run main.go YYYY XX` where `YYYY` is
the year of the event and `XX` is the day.
For instance, if you wish to run the solution for the 2nd day of the 2023 edition,
use `go run main.go 2023 02`.

This repo is still very much a work in progress and will regularly be updated
to include new puzzle solutions and refactoring as I learn more about the language.
