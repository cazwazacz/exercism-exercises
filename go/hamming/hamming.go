package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands must be of the same length")
	}

	return checkDifference(a, b), nil
}

func checkDifference(a, b string) int {
	difference := 0

	for i := range a {
		if a[i] != b[i] {
			difference++
		}
	}

	return difference
}
