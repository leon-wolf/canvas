package canvas

import (
	"image"
	"math"
)

func Resize(img image.Image, newWidth int, newHeight int, intFunc string) image.Image {

	if newWidth == 0 && newHeight == 0 {
		newWidth = img.Bounds().Dx()
		newHeight = img.Bounds().Dy()
	} else if newWidth == 0 && newHeight != 0 {
		newWidth = img.Bounds().Dx() / img.Bounds().Dy() * newHeight
	} else if newWidth != 0 && newHeight == 0 {
		newHeight = img.Bounds().Dy() / img.Bounds().Dx() * newWidth
	}
	// perform resizing
	switch intFunc {
	case "NearestNeighbor":
		return ScaleNearestNeighbor(img, newWidth, newHeight)
	default:
		return ScaleNearestNeighbor(img, newWidth, newHeight)
	}
}

func ScaleNearestNeighbor(img image.Image, newWidth int, newHeight int) image.Image {
	targetRect := image.Rect(0, 0, newWidth, newHeight)
	targetImg := image.NewNRGBA(targetRect)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			srcX := int(math.Round(float64(x) / float64(newWidth) * float64(width)))
			srcY := int(math.Round(float64(y) / float64(newHeight) * float64(height)))
			srcX = int(math.Min(float64(srcX), float64(width-1)))
			srcX = int(math.Min(float64(srcY), float64(height-1)))
			tarPixel := img.At(srcX, srcY)
			targetImg.Set(x, y, tarPixel)
		}
	}
	return targetImg
}
