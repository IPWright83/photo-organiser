package main

import "github.com/schollz/progressbar/v3"

var DefaultTheme = progressbar.Theme{
	Saucer:        "[green]=[reset]",
	SaucerHead:    "[green]>[reset]",
	SaucerPadding: " ",
	BarStart:      "[",
	BarEnd:        "]",
}
