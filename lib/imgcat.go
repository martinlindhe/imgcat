// Package imgcat provides helpers for embedding images (gif, png, jpeg)
// in the terminal output as suppored by iTerm 2.9+
// and documented at https://www.iterm2.com/images.html
package imgcat

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

// Cat embeds given io.Reader in the given io.Writer
func Cat(r io.Reader, w io.Writer) error {

	return embed(r, w)
}

// CatImage embeds given image.Image in the given io.Writer
func CatImage(i image.Image, w io.Writer) error {

	b, err := imageAsPngBytes(i)
	if err != nil {
		return err
	}
	return embed(b, w)
}

// CatFile embeds image file in the given io.Writer
func CatFile(fileName string, w io.Writer) error {

	r, err := os.Open(fileName)
	if err != nil {
		return err
	}

	return embed(r, w)
}

func embed(r io.Reader, w io.Writer) error {

	var inScreen bool
	switch os.Getenv("TERM") {
	case
		"screen",
		"tmux-256color":
		inScreen = true
	default:
		inScreen = false
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return err
	}

	// tmux requires unrecognized OSC sequences to be wrapped with DCS tmux;
	// <sequence> ST, and for all ESCs in <sequence> to be replaced with ESC ESC. It
	// only accepts ESC backslash for ST.
	if inScreen {
		fmt.Fprint(w, "\033Ptmux;\033")
	}
	fmt.Fprint(w, "\033]1337;File=;inline=1:")

	encoder := base64.NewEncoder(base64.StdEncoding, w)
	_, err = encoder.Write(buf.Bytes())
	if err != nil {
		return err
	}
	encoder.Close()

	fmt.Fprintln(w, "\a")
	if inScreen {
		fmt.Fprintln(w, "\033\\")
	}
	return nil
}

func imageAsPngBytes(i image.Image) (io.Reader, error) {

	buf := new(bytes.Buffer)
	err := png.Encode(buf, i)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
