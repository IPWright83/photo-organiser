package main

import (
	"errors"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"log"
	"os"
)

/**
 * This function grabs the year from the file camera data
 */
func getYearFromExif(filePath string) (int, error) {
	// Open the file, which is required for extracting
	// exif data (which is camera specific file attributes)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	exifData, err := exif.Decode(file)
	if err != nil {
		return 0, errors.New("No exif data avaliable")
	}

	time, err := exifData.DateTime()
	if err != nil {
		return 0, errors.New("No timestamp avaliable in exif data")
	}

	lat, long, _ := exifData.LatLong()
	fmt.Println(lat, long)

	return time.Year(), nil
}

/**
 * Obtain the most likely year that the photo was created
 * @return The year as a string
 */
func getYear(imageInfo ImageInfo) int {
	year, err := getYearFromExif(imageInfo.Path)
	if err != nil {
		year = imageInfo.Info.ModTime().Year()
	}

	return year
}
