// Package twofer contains a function which returns a sharing string.
package twofer

import "fmt"

// ShareWith takes a name and returns a "One for you, one for me." type string.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
