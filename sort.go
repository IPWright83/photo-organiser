package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func moveFileToYear(outputDir string, imageInfos []ImageInfo) {
	for _, imageInfo := range imageInfos {
		year := getYear(imageInfo)

		targetDir := outputDir + "/" + strconv.Itoa(year)
		log.Println(targetDir)
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			log.Println("making directory")
			os.Mkdir(targetDir, 0777) // Read & Write permission for everyone
		}

		targetPath := targetDir + "/" + imageInfo.Info.Name()
		err := os.Rename(imageInfo.Path, targetPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

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
	// fmt.Println(len(files))
	// fmt.Println(outputDirectory)
}
