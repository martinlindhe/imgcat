package main

import (
	"github.com/martinlindhe/imgcat"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file = kingpin.Arg("file", "A image file.").Required().File()
)

func main() {
	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	inFileName := (*file).Name()

	imgcat.CatImage(inFileName)
}
