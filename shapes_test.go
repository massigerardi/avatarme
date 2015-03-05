package avatarme

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"testing"
)

func TestDrawShape(t *testing.T) {
	fg := color.NRGBA{0x00, 0x00, 0xff, 0xff}
	bg := color.NRGBA{0xff, 0x00, 0x00, 0xff}
	for i, _ := range Shapes {
		img := DrawShape(bg, fg, i, 30)
		err := toFile(img, fmt.Sprintf("test/shapes/shape%v", i))
		if err != nil {
			t.Errorf("error: %v", err)
		}
	}
}

func TestDrawShapeHorMirror(t *testing.T) {
	fg := color.NRGBA{0x00, 0x00, 0xff, 0xff}
	bg := color.NRGBA{0xff, 0x00, 0x00, 0xff}
	for i, _ := range Shapes {
		img := DrawShapeHorMirror(bg, fg, i, 30)
		err := toFile(img, fmt.Sprintf("test/shapes/shape%vMirror", i))
		if err != nil {
			t.Errorf("error: %v", err)
		}
	}
}
func TestDrawShapeRotated(t *testing.T) {
	fg := color.NRGBA{0x00, 0x00, 0xff, 0xff}
	bg := color.NRGBA{0xff, 0x00, 0x00, 0xff}
	for i, _ := range Shapes {
		img := DrawShapeRotated(bg, fg, i, 30, 45*(math.Pi/180.0))
		err := toFile(img, fmt.Sprintf("test/shapes/shape%vRotated", i))
		if err != nil {
			t.Errorf("error: %v", err)
		}
	}
}

func toFile(img image.Image, p string) error {
	file, err := os.Create(fmt.Sprintf("%v%v", p, ".png"))
	if err != nil {
		return err
	}
	defer file.Close()
	png.Encode(file, img)
	return nil
}
