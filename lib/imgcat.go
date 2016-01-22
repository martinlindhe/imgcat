package imgcat

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
)

// CatRGBA embeds given image.RGBA in the terminal output (iTerm 2.9+)
func CatRGBA(i *image.RGBA) {

	data := imageAsPngBytes(i)

	embed(data)
}

// CatImage embeds given image.Image in the terminal output (iTerm 2.9+)
func CatImage(i *image.Image) {

	data := imageAsPngBytes(*i)

	embed(data)
}

// CatFile embeds given image file in the terminal output (iTerm 2.9+)
func CatFile(fileName string) error {

	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	embed(raw)
	return nil
}

func embed(data []byte) {
	enc := base64.StdEncoding.EncodeToString(data)

	printOsc()
	fmt.Print("1337;File=")
	fmt.Printf(";size=%d", len(enc))
	fmt.Print(";inline=1")
	fmt.Print(":")
	fmt.Print(enc)
	printSt()
	fmt.Println("")
}

func imageAsPngBytes(i image.Image) []byte {

	buf := new(bytes.Buffer)
	err := png.Encode(buf, i)
	if err != nil {
		fmt.Println(err)
	}
	return buf.Bytes()
}

// tmux requires unrecognized OSC sequences to be wrapped with DCS tmux;
// <sequence> ST, and for all ESCs in <sequence> to be replaced with ESC ESC. It
// only accepts ESC backslash for ST.
func printOsc() {
	fmt.Print("\033Ptmux;\033\033]")
}

// More of the tmux workaround described above.
func printSt() {
	fmt.Print("\a\033\\")
}
