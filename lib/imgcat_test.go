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

var (
	inFile = "../file.jpg"
)

func TestCat(t *testing.T) {

	f, err := os.Open(inFile)
	assert.Equal(t, nil, err)

	err = Cat(f, os.Stdout)
	assert.Equal(t, nil, err)
}

func TestCatFile(t *testing.T) {

	err := CatFile(inFile, os.Stdout)
	assert.Equal(t, nil, err)
}

func TestCatImage(t *testing.T) {

	img, err := decodeImage(inFile)
	assert.Equal(t, nil, err)

	err = CatImage(img, os.Stdout)
	assert.Equal(t, nil, err)
}

func TestCatRGBA(t *testing.T) {

	canvas := image.NewRGBA(image.Rect(0, 0, 20, 20))
	canvas.Set(10, 10, image.NewUniform(color.RGBA{255, 255, 255, 255}))

	err := CatImage(canvas, os.Stdout)
	assert.Equal(t, nil, err)
}

func decodeImage(filename string) (image.Image, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(bufio.NewReader(f))

	return img, err
}
