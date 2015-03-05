package main

import (
	"crypto/rand"
	"crypto/rsa"
	"flag"
	"fmt"
	"github.com/massigerardi/avatarme"
	"image"
	"image/png"
	"os"
)

var g avatarme.IconGenerator

func main() {
	var email = flag.String("email", "", "the email to encode")
	//var ip = flag.String("ip", "", "the ip to encode")
	//var key = flag.String("pKey", "", "the public key to use in the encoding")
	flag.Parse()
	pvtKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	if *email != "" {
		fmt.Printf("email %v\n", *email)
		g := avatarme.New280(pvtKey)
		img := g.GeneratePixel(*email)
		err = toFile(img, fmt.Sprintf("%v", *email))
		if err != nil {
			panic(err)
		}
	}

}

func toFile(img image.Image, p string) error {
	os.MkdirAll("demo/icons/", 0777)
	file, err := os.Create(fmt.Sprintf("demo/icons/%v.png", p))
	if err != nil {
		return err
	}
	fmt.Printf("write %v\n", file.Name())
	defer file.Close()
	png.Encode(file, img)
	fmt.Printf("wrote %v\n", file.Name())
	return nil
}
