package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

type ImageInfo struct {
	Path string
	Info fs.FileInfo
}

/**
 * Determine if the given file looks like an image
 */
func isImage(fileInfo fs.FileInfo) bool {
	extensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	extension := filepath.Ext(fileInfo.Name())
	return slices.Contains(extensions, extension)
}

/**
 * Returns the set of ImageInfos from the given directory, searching recursively
 */
func getImageInfos(root string) ([]ImageInfo, error) {
	var imageInfos []ImageInfo

	// err := filepath.Walk(root, func(path, string, info os.FileInfo, err error) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// We don't care about directories
		if !info.IsDir() {
			if isImage(info) {
				imageInfos = append(imageInfos, ImageInfo{Path: path, Info: info})
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return imageInfos, nil
}
