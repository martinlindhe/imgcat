# About

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
    imgcat.CatImage("file.jpg")
}
```


# License

Under [MIT](LICENSE)
