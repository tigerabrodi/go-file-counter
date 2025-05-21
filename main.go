package main

import (
	"bytes"
	"fmt"
	"os"
)

const statsFilename = "stats.txt"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program requires passing path to the file")
		os.Exit(1)
	}

	filename := os.Args[1]

	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Couldn't read file %s, %v", filename, err)
		os.Exit(1)
	}

	linesCount := countLines(data)
	wordsCount := countWords(data)

	content := fmt.Sprintf("Words: %d\n, lines: %d\n", wordsCount, linesCount)

	os.WriteFile(statsFilename, []byte(content), 0644)

}

func countLines(data []byte) int {
	return bytes.Count(data, []byte{'\n'}) + 1
}

func countWords(data []byte) int {
	return bytes.Count(data, []byte{' '}) + 1
}
