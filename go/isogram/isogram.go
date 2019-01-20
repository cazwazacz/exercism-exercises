package isogram

import (
	"unicode"
)

// IsIsogram takes a word and returns whether or not it is an isogram
func IsIsogram(word string) bool {
	record := make(map[rune]int)

	for _, character := range word {
		if character == '-' || character == ' ' {
			continue
		}

		character = unicode.ToLower(character)

		if record[character] > 0 {
			return false
		}

		record[character] = 1
	}

	return true
}
