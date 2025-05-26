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

	stats := analyzeText(data)

	content := fmt.Sprintf("Here are all your stats\n"+
		"Word count %d\n"+
		"Lines count: %d"+
		"Char count: %d"+
		"Avg word len %f",
		stats.wordsCount, stats.linesCount, stats.charsCount, stats.avgWordLen)

	os.WriteFile(statsFilename, []byte(content), 0644)

}

func countLines(data []byte) int {
	return bytes.Count(data, []byte{'\n'}) + 1
}

func countWords(data []byte) int {
	wordSlices := bytes.Fields(data)
	return len(wordSlices)
}

func countChars(data []byte) int {
	bytesWithoutWeirdWindowsBullshit := bytes.ReplaceAll(data, []byte{'\r'}, []byte{})
	return len(bytesWithoutWeirdWindowsBullshit)
}

type textStats struct {
	wordsCount int
	charsCount int
	linesCount int
	avgWordLen float64
}

func analyzeText(data []byte) textStats {
	stats := textStats{}

	stats.linesCount = countLines(data)
	stats.wordsCount = countWords(data)
	stats.charsCount = countChars(data)

	return stats
}
