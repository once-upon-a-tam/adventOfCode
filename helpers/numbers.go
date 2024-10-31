package helpers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)
var (
  blankStringRegexp = regexp.MustCompile(`^$`)
)

func IntsFromString(str, sep string) ([]int, error) {
	words := strings.Split(str, sep)

	ints := make([]int, 0)

	for _, w := range words {
    if blankStringRegexp.MatchString(w) {
      continue
    }
		n, err := strconv.Atoi(w)
		if err != nil {
			return nil, fmt.Errorf("%q is not an integer", w)
		}

    ints = append(ints, n)
	}

	return ints, nil
}
