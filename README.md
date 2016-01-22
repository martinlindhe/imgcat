# About

[![GoDoc](https://godoc.org/github.com/martinlindhe/imgcat?status.svg)](https://godoc.org/github.com/martinlindhe/imgcat)


Go port of the iTerm2 imgcat script

* https://www.iterm2.com/images.html
* https://raw.githubusercontent.com/gnachman/iTerm2/master/tests/imgcat

NOTE: this requires the use of iTerm 2.9 or later.


# Install command line

    go get -u github.com/martinlindhe/imgcat


# Use the lib in your terminal app

```go
package main

import "github.com/martinlindhe/imgcat/lib"

func main() {
    inFile := "file.jpg"

    // using a io.Reader
	f, _ := os.Open(inFile)
	imgcat.Cat(f, os.Stdout)

    // using filename
    imgcat.CatFile(inFile, os.Stdout)

    // using a image.RGBA
    canvas := image.NewRGBA(image.Rect(0, 0, 20, 20))
    canvas.Set(10, 10, image.NewUniform(color.RGBA{255, 255, 255, 255}))
    imgcat.CatRGBA(canvas, os.Stdout)
}
```


# License

Under [MIT](LICENSE)
