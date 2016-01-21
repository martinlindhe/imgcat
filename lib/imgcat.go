package imgcat

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

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

func CatImage(inFile string) {

	raw, err := ioutil.ReadFile(inFile)
	if err != nil {
		panic(err)
	}

	enc := base64.StdEncoding.EncodeToString(raw)

	printOsc()
	fmt.Print("1337;File=")
	fmt.Print("name=" + base64.StdEncoding.EncodeToString([]byte(inFile)) + ";")
	fmt.Printf("size=%d", len(enc))
	fmt.Print(";inline=1")
	fmt.Print(":")
	fmt.Print(enc)
	printSt()
	fmt.Println("")
}
