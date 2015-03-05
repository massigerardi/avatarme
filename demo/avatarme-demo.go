package main

import (
  "image"
  "image/png"
  "github.com/massigerardi/avatarme"
  "fmt"
  "os"
  "crypto/rsa"
  "crypto/rand"
)

var g avatarme.IconGenerator

func main() {
  fmt.Println("---------")
  pvtKey, err :=  rsa.GenerateKey(rand.Reader, 1024)
  if err != nil {
    panic(err)
  }
  g := avatarme.New280(pvtKey)
  img := g.GeneratePixel("massigerardi")
  toFile(img, "demo/icons/testenc")
  g = avatarme.New400(pvtKey)
  img = g.GeneratePixel("massigerardi")
  toFile(img, "demo/icons/test400enc")
  g = avatarme.New280(nil)
  img = g.GeneratePixel("massigerardi")
  toFile(img, "demo/icons/test250")
  g = avatarme.New400(nil)
  img = g.GeneratePixel("massigerardi")
  toFile(img, "demo/icons/test400")
}

func toFile(img image.Image, p string) error {
  file, err := os.Create(fmt.Sprintf("%v.png", p))
  if err != nil {
    return err
  }
  defer file.Close()
  png.Encode(file, img)
  return nil
}