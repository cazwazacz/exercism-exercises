package isogram

import (
	"unicode"
)

var record = make(map[rune]int)

// IsIsogram takes a word and returns whether or not it is an isogram
func IsIsogram(word string) bool {
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
