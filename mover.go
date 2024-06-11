package main

import (
	"fmt"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func modifyFileName(targetDir string, imageInfo ImageInfo) string {
	fileName := imageInfo.Info.Name()
	extension := filepath.Ext(fileName)
	name := fileName[:len(fileName)-len(extension)]

	// Create the target path
	targetPath := targetDir + "/" + name + strings.ToLower(extension)
	targetFile := targetPath
	counter := 0

	// Ensure that we don't overwrite existing files by checking for their
	// presence first and adding a counter onto the end to prevent a clash
	for {
		if _, err := os.Stat(targetFile); os.IsNotExist(err) {
			return targetFile
		}

		counter++
		targetFile = targetDir + "/" + fmt.Sprintf("%s_%d%s", name, counter, strings.ToLower(extension))
	}
}

func moveFileToYear(outputDir string, imageInfos []ImageInfo) {
	bar := progressbar.NewOptions(len(imageInfos),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetDescription("Moving files"),
		progressbar.OptionSetTheme(DefaultTheme),
	)

	for _, imageInfo := range imageInfos {
		year := getYear(imageInfo)

		targetDir := outputDir + "/" + strconv.Itoa(year)

		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			//log.Println("making directory")
			os.Mkdir(targetDir, 0777) // Read & Write permission for everyone
		}

		targetPath := modifyFileName(targetDir, imageInfo)

		fmt.Printf("Moving %s to %s\n", imageInfo.Info.Name(), targetPath)

		err := os.Rename(imageInfo.Path, targetPath)
		if err != nil {
			log.Fatal(err)
		}
		bar.Add(1)
	}

	bar.Finish()
}
