package main

import (
	"fmt"
	"os"

	"gifgenerator"
)

func main() {
	str := "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	gifgen := gifgenerator.NewGIFGenerator(128, 128, 50)
	gif := gifgen.EncodeGif(str[32*2:])
	gifFile, err := os.OpenFile("hello.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gocat: %v\n", err)
		os.Exit(1)
	}
	defer gifFile.Close()

	gifFile.Write(gif.Bytes())
}
