package main

import (
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
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
	extensions := []string{".jpg", ".jpeg", ".png", ".gif", ".JPG", ".JPEG", ".PNG", ".GIF"}
	extension := filepath.Ext(fileInfo.Name())
	return slices.Contains(extensions, extension)
}

/**
 * Returns the set of ImageInfos from the given directory, searching recursively
 */
func getImageInfos(root string) ([]ImageInfo, error) {
	var imageInfos []ImageInfo

	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetDescription("Scanning for files"),
		progressbar.OptionSetTheme(DefaultTheme),
	)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// We don't care about directories
		if !info.IsDir() {
			bar.Add(1)
			if isImage(info) {
				imageInfos = append(imageInfos, ImageInfo{Path: path, Info: info})
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	bar.Finish()
	return imageInfos, nil
}
