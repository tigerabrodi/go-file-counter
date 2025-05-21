package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program requires passing path to the file")
		os.Exit(1)
	}

	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		log.Printf("Error, couldnt open file %s, some error: %v", filename, err)

		fmt.Fprintf(os.Stderr, "Something went wrong when opening the file %s", filename)
		fmt.Fprintf(os.Stderr, "Make sure the file exists so that we can open it, maybe the file name itself is wrong?")
	}

	defer file.Close()
}
