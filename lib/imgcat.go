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

// CatRGBA embeds given image.RGBA in the terminal output (iTerm 2.9+)
func CatRGBA(i *image.RGBA) error {

	return embed(imageAsPngBytes(i), os.Stdout)
}

// CatImage embeds given image.Image in the terminal output (iTerm 2.9+)
func CatImage(i *image.Image) error {

	return embed(imageAsPngBytes(*i), os.Stdout)
}

// CatFile embeds given image file in the terminal output (iTerm 2.9+)
func CatFile(fileName string) error {

	r, err := os.Open(fileName)
	if err != nil {
		return err
	}

	embed(r, os.Stdout)
	return nil
}

// CatReader embeds given io.Reader in the given io.Writer
func CatReader(r io.Reader, w io.Writer) error {

	return embed(r, w)
}

func embed(r io.Reader, w io.Writer) error {

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	fmt.Printf("wrote %d bytes\n", len(buf.Bytes()))

	// tmux requires unrecognized OSC sequences to be wrapped with DCS tmux;
	// <sequence> ST, and for all ESCs in <sequence> to be replaced with ESC ESC. It
	// only accepts ESC backslash for ST.
	fmt.Fprint(w, "\033Ptmux;\033\033]")

	fmt.Fprintf(w, "1337;File=;inline=1:")

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

func imageAsPngBytes(i image.Image) io.Reader {

	buf := new(bytes.Buffer)
	err := png.Encode(buf, i)
	if err != nil {
		fmt.Println(err)
	}
	return buf
}
