// This file serves as the entry point for the program
// Although not a requirement, by convention, the main package is usually defined in a file named main.go
package main

import (
	"bufio" // To read text
	"flag"  // Create and manage command-line flags
	"fmt"   // To print formatted output
	"io"    // Provides the io.Reader interface
	"os"    // Use os resources
)

func count(r io.Reader, countLines bool) int {
	// A scanner is used to read text from a Reader(such as files)
	scanner := bufio.NewScanner(r)

	// If the count lines flag is not set, we want to count words so we define
	// the scanner split type to words(default is split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	// Defining a counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Parsing the flags provided by the user
	flag.Parse()
	// Calling the count function to count the number of words (or lines)
	// received from the Standard Input and printing out
	fmt.Println(count(os.Stdin, *lines))
}
