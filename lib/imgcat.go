package imgcat

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"
)

// CatRGBA embeds given image.RGBA in the terminal output (iTerm 2.9+)
func CatRGBA(i *image.RGBA) {

	data := imageAsPngBytes(i)

	embed(data, os.Stdout)
}

// CatImage embeds given image.Image in the terminal output (iTerm 2.9+)
func CatImage(i *image.Image) {

	data := imageAsPngBytes(*i)

	embed(data, os.Stdout)
}

// CatFile embeds given image file in the terminal output (iTerm 2.9+)
func CatFile(fileName string) error {

	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	embed(raw, os.Stdout)
	return nil
}

// CatReader embeds given io.Reader in the given io.Writer
func CatReader(r io.Reader, w io.Writer) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	embed(buf.Bytes(), w)
}

func embed(data []byte, w io.Writer) {
	enc := base64.StdEncoding.EncodeToString(data)

	// tmux requires unrecognized OSC sequences to be wrapped with DCS tmux;
	// <sequence> ST, and for all ESCs in <sequence> to be replaced with ESC ESC. It
	// only accepts ESC backslash for ST.
	fmt.Fprint(w, "\033Ptmux;\033\033]")

	fmt.Fprint(w, "1337;File=")
	fmt.Fprintf(w, ";size=%d", len(enc))
	fmt.Fprint(w, ";inline=1")
	fmt.Fprint(w, ":")
	fmt.Fprint(w, enc)

	// More of the tmux workaround described above.
	fmt.Println(w, "\a\033\\")
}

func imageAsPngBytes(i image.Image) []byte {

	buf := new(bytes.Buffer)
	err := png.Encode(buf, i)
	if err != nil {
		fmt.Println(err)
	}
	return buf.Bytes()
}
