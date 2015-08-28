package stringutil

// Package stringutil contains utility functions for working with strings.
// Reverse returns its argument string reversed rune-wise left to right.

func Reverse(s string) string {
	return reverseTwo(s)
}

/*
go build
 won't produce an output file.

go install
 will place the package inside the pkg directory of the workspace.
*/
