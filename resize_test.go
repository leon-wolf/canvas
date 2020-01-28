package canvas

import (
	"image"
	"image/color"
	"log"
	"runtime"
	"testing"
)

var img = image.NewNRGBA(image.Rect(0, 0, 5, 5))

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	img.Set(5, 5, color.Black)
}

func Test_SameSize(t *testing.T) {
	newImg, _ := Resize(img, 0, 0, CatmullRom)
	if newImg.Bounds() != img.Bounds() {
		log.Print(img.Bounds().Dx(), "x", img.Bounds().Dy())
		log.Print(newImg.Bounds().Dx(), "x", newImg.Bounds().Dy())
		t.Fail()
	}
}

func Test_Height(t *testing.T) {
	newImg, err := Resize(img, 0, 100, CatmullRom)
	if err != nil || newImg.Bounds().Dy() != 100 || newImg.Bounds().Dx() != 100 {
		t.Fail()
	}
}
func Test_Width(t *testing.T) {
	newImg, err := Resize(img, 200, 0, CatmullRom)
	if err != nil || newImg.Bounds().Dy() != 200 || newImg.Bounds().Dx() != 200 {
		t.Fail()
	}
}
func Test_WidthAndHeight(t *testing.T) {
	newImg, err := Resize(img, 300, 300, CatmullRom)
	if err != nil || newImg.Bounds() != image.Rect(0, 0, 300, 300) {
		t.Fail()
	}
}
