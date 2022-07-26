package models

import (
	"gifgenerator"
	"os"
	"strconv"

	"github.com/abbb03/textgif/app/inputprocessor"
)

func CreateGif(text, sizeX, sizeY string) (*os.File, error) {
	text, err := inputprocessor.GetText(text)
	if err != nil {
		return nil, err
	}

	x, err := strconv.Atoi(sizeX)
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(sizeY)
	if err != nil {
		return nil, err
	}

	gifgen := gifgenerator.NewGIFGenerator(x, y, 50)
	gif := gifgen.EncodeGif(text)
	gifFile, err := os.OpenFile("test.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	defer gifFile.Close()

	gifFile.Write(gif.Bytes())

	return gifFile, nil
}
