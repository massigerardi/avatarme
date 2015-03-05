package avatarme

import (
	"testing"
)

var g IconGenerator

func TestGenerate(t *testing.T) {
	g := New250(nil)
	img := g.GenerateShapes("massigerardi")
	SaveToPng(img, "test/icons/shape250")
}
