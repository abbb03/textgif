package drawer

import (
	"image"
	"image/color"
)

func DrawVertical(img *image.Paletted, x, y, length int) {
	for i := 0; i < length; i++ {
		img.Set(x, y+i, color.Black)
	}
}

func DrawHorizontal(img *image.Paletted, x, y, length int) {
	for i := length; i > 0; i-- {
		img.Set(x+i, y, color.Black)
	}
}

func DrawFastDiagonal(img *image.Paletted, x, y, width int) {
	for i := 0; i < width; i++ {
		img.Set(x+i, y+i, color.Black)
	}
}

func DrawSlowDiagonal(img *image.Paletted, x, y, width int) {
	for i := 0; i < width; i++ {
		img.Set(x+i, y+width-i, color.Black)
	}
}

func DrawWhitespace(img *image.Paletted, x, y int) {
	img.Set(x, y, color.White)
}
