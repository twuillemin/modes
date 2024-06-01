package common

import "strings"

// Printable is the basic interface that most of the fields and messages should implement
type Printable interface {
	// ToString returns a basic, but readable, representation of the field
	ToString() string
}

// PrefixMultiLine is a helper functions, that given a string with line breaks will prefix each line with the given
// prefix
//
// Params:
//   - lines: The string containing the lines to split
//   - prefix: The prefix to apply to each line
//
// Returns the same string as the one given but prefixed
func PrefixMultiLine(lines string, prefix string) string {

	subLines := strings.Split(lines, "\n")

	newLines := make([]string, len(subLines))

	for i, line := range subLines {
		newLines[i] = prefix + line
	}

	return strings.Join(newLines, "\n")
}
