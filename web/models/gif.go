package models

import (
	"bytes"
	"gifgenerator"
	"os"
	"strconv"

	"github.com/abbb03/textgif/app/inputprocessor"
	"github.com/abbb03/textgif/web/pkg/namegen"
)

type gif struct {
	Name    string
	content bytes.Buffer
}

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
	g := &gif{
		content: gifgen.EncodeGif(text),
		Name:    namegen.GetMD5Hash(),
	}

	gifFile, err := os.OpenFile(g.Name+".gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	defer gifFile.Close()

	gifFile.Write(g.content.Bytes())

	return gifFile, nil
}
