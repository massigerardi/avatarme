package avatarme

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type IconGenerator interface {
	GeneratePixel(s string) image.Image
	GenerateShapes(s string) image.Image
}

type identicon struct {
	size       int
	pixSize    int
	margin     int
	rows, cols int
	privateKey *rsa.PrivateKey
	encrypted  bool
}

func New280(pvtKey *rsa.PrivateKey) IconGenerator {
	return newIdenticon(280, 30, 20, pvtKey)
}

func newIdenticon(size, pixSize, margin int, pvtKey *rsa.PrivateKey) IconGenerator {
	return &identicon{
		size:       size,
		pixSize:    pixSize,
		margin:     margin,
		rows:       (size - 2*margin) / pixSize,
		cols:       (size - 2*margin) / pixSize,
		privateKey: pvtKey,
		encrypted:  pvtKey != nil,
	}
}

func New400(pvtKey *rsa.PrivateKey) IconGenerator {
	return newIdenticon(400, 40, 20, pvtKey)
}

func (id identicon) GeneratePixel(s string) image.Image {
	b := id.gethash(s)
//	fmt.Printf("%v -> %x\n", s, b)
	return id.renderPixel(b)
}

func (id identicon) GenerateShapes(s string) image.Image {
	b := id.gethash(s)
//	fmt.Printf("%v -> %v\n", s, b)
	return id.renderShapes(b)
}

func (id identicon) gethash(s string) [28]byte {
	if id.encrypted {
		md5hash := md5.New()
		label := []byte("identicon")
		encryptedmsg, err := rsa.EncryptOAEP(md5hash, rand.Reader, &id.privateKey.PublicKey, []byte(s), label)
		if err != nil {
			panic(err)
		}
		return sha256.Sum224(encryptedmsg)
	} else {
		return sha256.Sum224([]byte(s))
	}
}

func (id identicon) renderPixel(b [28]byte) image.Image {
	bg := color.NRGBA{0xf0, 0xf0, 0xf0, 0xff}
	fg := color.NRGBA{uint8(b[0] << 2), uint8(b[1] << 4), uint8(b[3] << 8), 0xff}
	img := image.NewPaletted(image.Rect(0, 0, id.size, id.size), color.Palette{bg, fg})
	pixels := make([]byte, id.pixSize)
	for i := 0; i < id.pixSize; i++ {
		pixels[i] = 1
	}
	cell := 0
	for c := 0; c < id.cols; c++ {
		for r := 0; r < (id.rows+1)/2; r++ {
			//fmt.Printf("%v:%v --> %v\n",r,c, cell)
			index := cell % 28
//			fmt.Printf(" %v=%v[%x]:", cell, b[index], b[index])
			if b[index]%2 == 0 {
				x := id.margin + r*id.pixSize
				x1 := id.size - id.margin - r*id.pixSize - id.pixSize
				for i := 0; i < id.pixSize; i++ {
					y := id.margin + i + 1 + c*id.pixSize
					//fmt.Printf("%v:%v %v:%v\n",x,y,x1,y)
					offs := img.PixOffset(x, y)
					copy(img.Pix[offs:], pixels)
					offs = img.PixOffset(x1, y)
					copy(img.Pix[offs:], pixels)
				}
			}
			cell++
		}
//		fmt.Println()
	}
	return img
}

func (id identicon) renderShapes(b [28]byte) image.Image {
	bg := color.NRGBA{0xf0, 0xf0, 0xf0, 0xff}
	fg := color.NRGBA{uint8(b[0] << 2), uint8(b[0] << 4), uint8(b[0] << 8), 0xff}
	img := image.NewRGBA(image.Rect(0, 0, id.size, id.size))
	draw.Draw(img, img.Bounds(), &image.Uniform{bg}, image.ZP, draw.Src)
	cell := 0
//	fmt.Printf("%v * %v\nr:c\n", id.rows, id.cols)
	for c := 0; c < id.cols/2; c++ {
		for r := 0; r < id.rows; r++ {
			x := id.margin + c*id.pixSize
			x1 := id.size - id.margin - c*id.pixSize - id.pixSize
			y := id.margin + r*id.pixSize
			index := int(b[cell%len(b)])
			fg = color.NRGBA{uint8(index << 2), uint8(index << 4), uint8(index << 8), 0xff}
			//fmt.Printf("%v:%v -> %v:%v %v:%v\n", r, c, x, y, x1, y)
			DrawShapeOnImage(fg, fg, index, id.pixSize, img, image.Pt(x, y), 0.0)
			DrawShapeHorMirrorOnImage(fg, fg, index, id.pixSize, img, image.Pt(x1, y))
			cell++
//			fmt.Println()
		}
	}
	return img
}

func SaveToPng(img image.Image, p string) error {
	file, err := os.Create(fmt.Sprintf("%v%v", p, ".png"))
	if err != nil {
		return err
	}
	defer file.Close()
	png.Encode(file, img)
	return nil
}
