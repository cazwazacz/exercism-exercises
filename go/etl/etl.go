package etl

import "strings"

// Transform takes a legacy form of storing scrabble scores and returns a more efficient structure
func Transform(legacy map[int][]string) map[string]int {
	var transformed = make(map[string]int)

	for points, letters := range legacy {
		for _, letter := range letters {
			transformed[strings.ToLower(letter)] = points
		}
	}

	return transformed
}
