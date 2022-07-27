package models

import (
	"gifgenerator"
	"os"
	"strconv"

	"github.com/abbb03/textgif/app/inputprocessor"
	"github.com/abbb03/textgif/web/pkg/namegen"
)

type gif struct {
	Path string
	File *os.File
}

func GetGif(text, sizeX, sizeY string) (*gif, error) {
	giffile, err := createGifFile(text, sizeX, sizeY)
	if err != nil {
		return &gif{}, err
	}
	g := &gif{}
	g.Path = "../../../" + giffile.Name()
	g.File = giffile
	return g, nil
}

func createGifFile(text, sizeX, sizeY string) (*os.File, error) {
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
	content := gifgen.EncodeGif(text)
	filename := namegen.GetName()

	gifFile, err := os.OpenFile(filename+".gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}

	defer gifFile.Close()

	gifFile.Write(content.Bytes())

	return gifFile, nil
}
