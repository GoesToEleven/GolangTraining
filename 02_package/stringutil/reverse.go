// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	return reverseTwo(s)
}

/*
go build
	go build reverse.go reverseTwo.go
 	won't produce an output file.

go install
 	will place the package inside the pkg directory of the workspace.
*/
