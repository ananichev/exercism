// Package twofer implements simple ShareWith function.
package twofer

import "fmt"

// ShareWith return a string, "name = "you"" if name is passed,
// "One for you, one for me." if name is blank
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
