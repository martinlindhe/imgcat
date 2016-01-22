package imgcat

import (
	"bufio"
	"image"
	"image/color"
	_ "image/jpeg"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatFile(t *testing.T) {

	inFile := "file.jpg"
	CatFile(inFile)
}

func TestCatImage(t *testing.T) {

	inFile := "file.jpg"

	img, err := decodeImage(inFile)
	assert.Equal(t, nil, err)

	CatImage(&img)
}

func TestCatRGBA(t *testing.T) {

	canvas := image.NewRGBA(image.Rect(0, 0, 20, 20))
	canvas.Set(10, 10, image.NewUniform(color.RGBA{255, 255, 255, 255}))

	CatRGBA(canvas)
}

func TestCatReader(t *testing.T) {

	inFile := "file.jpg"

	f, err := os.Open(inFile)
	assert.Equal(t, nil, err)

	CatReader(f, os.Stdout)
}

// returns image.Image, mime-type string, error
func decodeImage(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(bufio.NewReader(f))

	return img, err
}
