package models

import (
	"gifgenerator"
	"os"
	"path/filepath"
	"strconv"

	"github.com/abbb03/textgif/app/inputprocessor"
	"github.com/abbb03/textgif/web/pkg/namegen"
)

func CreateGifFile(text, sizeX, sizeY string) (*os.File, error) {
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
	filename := namegen.GetName() + ".gif"
	path := filepath.Join("./tmp", filename)
	gifFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}

	defer gifFile.Close()

	gifFile.Write(content.Bytes())

	return gifFile, nil
}
