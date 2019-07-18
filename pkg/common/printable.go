package common

import "strings"

// Printable is the basic interface that most of fields and messages should implement
type Printable interface {
	// ToString returns a basic, but readable, representation of the field
	ToString() string
}

func PrefixMultiLine(lines string, prefix string) string {

	subLines := strings.Split(lines, "\n")

	newLines := make([]string, len(subLines))

	for i, line := range subLines {
		newLines[i] = prefix + line
	}

	return strings.Join(newLines, "\n")
}
