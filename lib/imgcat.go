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

// CatRGBA embeds given image.RGBA in the given io.Writer
func CatRGBA(i *image.RGBA, w io.Writer) error {

	b, err := imageAsPngBytes(i)
	if err != nil {
		return err
	}
	return embed(b, w)
}

// CatImage embeds given image.Image in the given io.Writer
func CatImage(i *image.Image, w io.Writer) error {

	b, err := imageAsPngBytes(*i)
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

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	// tmux requires unrecognized OSC sequences to be wrapped with DCS tmux;
	// <sequence> ST, and for all ESCs in <sequence> to be replaced with ESC ESC. It
	// only accepts ESC backslash for ST.
	fmt.Fprint(w, "\033Ptmux;\033\033]1337;File=;inline=1:")

	encoder := base64.NewEncoder(base64.StdEncoding, w)
	defer encoder.Close()

	_, err := encoder.Write(buf.Bytes())
	if err != nil {
		return err
	}

	// More of the tmux workaround described above.
	fmt.Fprintln(w, "\a\033\\")
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
