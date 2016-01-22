package main

import (
	"fmt"
	"os"

	"github.com/martinlindhe/imgcat/lib"

	"gopkg.in/alecthomas/kingpin.v2"
)

type fileList []string

// exists reports whether the named file or directory exists.
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (i *fileList) Set(value string) error {

	if !exists(value) {
		return fmt.Errorf("'%s' does not exist", value)
	}
	*i = append(*i, value)
	return nil
}

func (i *fileList) String() string {
	return ""
}

func (i *fileList) IsCumulative() bool {
	return true
}

func imageList(s kingpin.Settings) (target *[]string) {
	target = new([]string)
	s.SetValue((*fileList)(target))
	return
}

var (
	files   = imageList(kingpin.Arg("files", "Image files to show.").Required())
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	heading := false
	if len(*files) > 1 && *verbose {
		heading = true
	}

	for _, file := range *files {
		if heading {
			fmt.Printf("%s:\n", file)
		}
		err := imgcat.CatFile(file, os.Stdout)
		if err != nil {
			fmt.Println(err)
		}
	}
}
