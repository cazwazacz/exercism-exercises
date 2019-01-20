package isogram

import (
	"unicode"
)

// IsIsogram takes a word and returns whether or not it is an isogram
func IsIsogram(word string) bool {
	record := make(map[rune]bool)

	for _, character := range word {
		if character == '-' || character == ' ' {
			continue
		}

		character = unicode.ToLower(character)

		if record[character] {
			return false
		}

		record[character] = true
	}

	return true
}
