package main

import (
	"fmt"
	"os"

	"gifgenerator"

	"github.com/abbb03/textgif/app/inputprocessor"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "textgif: enter the filename\n")
		return
	}
	gifName := os.Args[1]
	text, err := inputprocessor.GetInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	gifgen := gifgenerator.NewGIFGenerator(96, 96, 50)
	gif := gifgen.EncodeGif(text)
	gifFile, err := os.OpenFile(gifName+".gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gocat: %v\n", err)
		return
	}
	defer gifFile.Close()

	gifFile.Write(gif.Bytes())
}
