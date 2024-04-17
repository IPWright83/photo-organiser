package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Ensure the program has been called correctly first
	if len(os.Args) < 3 {
		fmt.Println("You need to call with {inputDirectory} and {outputDirectory}")
		return
	}

	sourceDir := os.Args[1]
	outputDir := os.Args[2]

	// Grab all the files within the source directory
	imageFiles, err := getImageInfos(sourceDir)
	if err != nil {
		log.Fatal(err)
	}

	fileCount := len(imageFiles)
	fmt.Printf("Discovered %d files\n", fileCount)

	moveFileToYear(outputDir, imageFiles)
	fmt.Printf("Moved %d files\n", fileCount)
}
