// This file serves as the entry point for the program
// Although not a requirement, by convention, the main package is usually defined in a file named main.go
package main

import (
	"bufio" // To read text
	"fmt"   // To print formatted output
	"io"    // Provides the io.Reader interface
	"os"    // Use os resources
)

func count(r io.Reader) int {
	// A scanner is used to read text from a Reader(such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words(default is split by lines)
	scanner.Split(bufio.ScanWords)

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
	// Calling the count function to count the number of words
	// received from the Standard Input and printing out
	fmt.Println(count(os.Stdin))
}
