// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {

	}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "%v", err)
		}
	}

}
func countFilesSameLine(f *os.File, counts map[string]int) {

}
