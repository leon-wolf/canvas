package canvas

import (
	"golang.org/x/image/draw"
	"image"
)

type InterpolationFunction int

// InterpolationFunction constants
const (
	Bilinear InterpolationFunction = iota
	CatmullRom
)

// kernal, returns an InterpolationFunctions taps and kernel.
func (i InterpolationFunction) kernel() *draw.Kernel {
	switch i {
	case Bilinear:
		return draw.BiLinear
	case CatmullRom:
		return draw.CatmullRom
	default:
		return draw.CatmullRom
	}
}

func Resize(img image.Image, width int, height int, intFunc InterpolationFunction) (image.Image, error) {

	if width == 0 && height == 0 {
		width = img.Bounds().Dx()
		height = img.Bounds().Dy()
	} else if width == 0 && height != 0 {
		width = img.Bounds().Dx() / img.Bounds().Dy() * height
	} else if width != 0 && height == 0 {
		height = img.Bounds().Dy() / img.Bounds().Dx() * width
	}
	// new size of image
	dstRect := image.Rect(0, 0, width, height)
	// perform resizing
	dstImg := image.NewRGBA(dstRect)
	intFunc.kernel().Scale(dstImg, dstRect, img, img.Bounds(), draw.Over, nil)

	return dstImg, nil
}
