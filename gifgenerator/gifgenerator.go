package gifgenerator

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"

	"gifgenerator/drawer"
)

const (
	charX = 5
	charY = 7
)

type sheet struct {
	x     int
	y     int
	image *image.Paletted
}

type GIFGenerator struct {
	Sheet  *sheet
	Delay  int
	images []*image.Paletted
	delays []int
}

func NewGIFGenerator(sizeX, sizeY, delay int) *GIFGenerator {
	g := &GIFGenerator{
		Sheet: newSheet(sizeX, sizeY),
		Delay: delay,
	}

	return g
}

func (g *GIFGenerator) EncodeGif(text string) bytes.Buffer {
	g.images = append(g.images, g.Sheet.image)
	g.delays = append(g.delays, g.Delay)

	for _, val := range text {
		newframe := g.writeLetter(val)
		g.images = append(g.images, newframe)
		g.delays = append(g.delays, g.Delay)
	}

	gf := gif.GIF{
		Image: g.images,
		Delay: g.delays,
	}

	var buf bytes.Buffer
	gif.EncodeAll(&buf, &gf)
	return buf
}

func newSheet(x, y int) *sheet {
	palette := []color.Color{color.White, color.Black}
	rect := image.Rect(0, 0, x, y)
	imgsheet := sheet{
		x:     1,
		y:     1,
		image: image.NewPaletted(rect, palette),
	}

	return &imgsheet
}

func (g *GIFGenerator) writeLetter(l rune) *image.Paletted {
	prevFrame := g.images[len(g.images)-1]
	frame := image.NewPaletted(prevFrame.Rect, prevFrame.Palette)
	copy(frame.Pix, prevFrame.Pix)
	switch l {
	case 'а':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y, charY)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y, charX-1)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y+1, charY-1)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y+3, charX-1)
	case 'б':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y, charY)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y, charX-1)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y+4, charY-5)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y+3, charX-1)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y+6, charX-1)
	case 'в':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y, charY)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y, charX-1)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y+charY-1, charX-1)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y+charY-4, charX-1)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y+1, charY-5)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y+4, charY-5)
	case 'г':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y, charY)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y, charX)
	case 'д':
		drawer.DrawVertical(frame, g.Sheet.x+1, g.Sheet.y, charY-2)
		drawer.DrawVertical(frame, g.Sheet.x+charX-1, g.Sheet.y+1, charY-3)
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y+5, 2)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y+5, 2)
		drawer.DrawHorizontal(frame, g.Sheet.x-1, g.Sheet.y+4, charX+1)
		drawer.DrawHorizontal(frame, g.Sheet.x+1, g.Sheet.y, charX-3)
	case 'е':
		fallthrough
	case 'ё':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y, charY)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y, charX-1)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y+3, charX)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y+charY-1, charX)
	case 'ж':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y, 2)
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y+charY-2, 2)
		drawer.DrawVertical(frame, g.Sheet.x+1, g.Sheet.y+2, 1)
		drawer.DrawVertical(frame, g.Sheet.x+1, g.Sheet.y+4, 1)
		drawer.DrawVertical(frame, g.Sheet.x+2, g.Sheet.y, charY)
		drawer.DrawVertical(frame, g.Sheet.x+4, g.Sheet.y, 2)
		drawer.DrawVertical(frame, g.Sheet.x+4, g.Sheet.y+charY-2, 2)
		drawer.DrawVertical(frame, g.Sheet.x+3, g.Sheet.y+2, 1)
		drawer.DrawVertical(frame, g.Sheet.x+3, g.Sheet.y+4, 1)
	case 'з':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y+1, 1)
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y+charY-2, 1)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y, charX-2)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y+charY-1, charX-2)
		drawer.DrawHorizontal(frame, g.Sheet.x+1, g.Sheet.y+charY-4, charX-3)
		drawer.DrawVertical(frame, g.Sheet.x+charX-1, g.Sheet.y+1, 2)
		drawer.DrawVertical(frame, g.Sheet.x+charX-1, g.Sheet.y+charY-3, 2)
	case 'и':
		fallthrough
	case 'й':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y, charY)
		drawer.DrawSlowDiagonal(frame, g.Sheet.x+1, g.Sheet.y, charX-1)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y, charY)
	case 'к':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y, charY)
		drawer.DrawSlowDiagonal(frame, g.Sheet.x+1, g.Sheet.y-1, charX-1)
		drawer.DrawFastDiagonal(frame, g.Sheet.x+1, g.Sheet.y+charY/2, charX-1)
	case 'л':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y+charY-1, 1)
		drawer.DrawVertical(frame, g.Sheet.x+1, g.Sheet.y+charY-1, 1)
		drawer.DrawVertical(frame, g.Sheet.x+2, g.Sheet.y, charY-1)
		drawer.DrawHorizontal(frame, g.Sheet.x+2, g.Sheet.y, charX-3)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y, charY)
	case 'м':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y, charY)
		drawer.DrawFastDiagonal(frame, g.Sheet.x+1, g.Sheet.y+charY/2-1, 2)
		drawer.DrawSlowDiagonal(frame, g.Sheet.x+2, g.Sheet.y+charY/2-2, 2)
		drawer.DrawVertical(frame, g.Sheet.x+charX-1, g.Sheet.y, charY)
	case 'н':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y, charY)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y+charY/2, charX)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y, charY)
	case 'о':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y+1, charY-2)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y, charX-1)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y+charY-1, charX-1)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y+1, charY-2)
	case 'п':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y, charY)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y, charX)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y, charY)
	case 'р':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y, charY)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y, charX-1)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y+3, charX-1)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y+1, 2)
	case 'с':
		drawer.DrawVertical(frame, g.Sheet.x, g.Sheet.y+1, charY-2)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y, charX-1)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y+charY-1, charX-1)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y+1, 1)
		drawer.DrawVertical(frame, g.Sheet.x+charX, g.Sheet.y+charY-2, 1)
	case 'т':
		drawer.DrawVertical(frame, g.Sheet.x+charX/2+1, g.Sheet.y+1, charY-1)
		drawer.DrawHorizontal(frame, g.Sheet.x, g.Sheet.y, charX)
	case 'у':
		drawer.DrawFastDiagonal(frame, g.Sheet.x, g.Sheet.y, 3)
		drawer.DrawSlowDiagonal(frame, g.Sheet.x+3, g.Sheet.y-1, 2)
	default:
		return frame
	}

	g.Sheet.x += charX + 2
	g.Sheet.y = 1

	return frame
}
