package proverb

import "fmt"

// Proverb takes strings and returns a proverb made from the.
func Proverb(rhyme []string) []string {
	var proverb []string
	if len(rhyme) > 0 {
		if len(rhyme) > 1 {
			for i := 0; i < len(rhyme)-1; i++ {
				proverb = append(proverb, fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1]))
			}
		}

		proverb = append(proverb, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
	}

	return proverb
}
