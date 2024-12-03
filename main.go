package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/nfnt/resize"
)

func main() {
	reszingimg := resizeJpg()
	toWebp(reszingimg, "test-img/image-jpg.webp")

	reszingimg = resizePng()
	toWebp(reszingimg, "test-img/image-png.webp")
}

func resizeJpg() image.Image {
	// Load image from file
	imageDataJpg, err := os.ReadFile("test-img/image.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	// Decode image
	img, _, err := image.Decode(bytes.NewReader(imageDataJpg))
	if err != nil {
		log.Fatalln(err)
	}

	// Resize image
	reszingimg := resize.Resize(360, 0, img, resize.Lanczos3) // width 360px, height auto
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, reszingimg, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		log.Fatalln(err)
	}

	return reszingimg
}

func resizePng() image.Image {
	// Load image from file
	imageDataPng, err := os.ReadFile("test-img/image.png")
	if err != nil {
		log.Fatalln(err)
	}

	// Decode image
	img, _, err := image.Decode(bytes.NewReader(imageDataPng))
	if err != nil {
		log.Fatalln(err)
	}

	// Resize image
	reszingimg := resize.Resize(360, 0, img, resize.Lanczos3) // width 360px, height auto
	buf := new(bytes.Buffer)
	err = png.Encode(buf, reszingimg)
	if err != nil {
		log.Fatalln(err)
	}

	return reszingimg
}

func toWebp(reszingimg image.Image, path string) {

	output, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Fatalln(err)
	}

	if err := webp.Encode(output, reszingimg, options); err != nil {
		log.Fatalln(err)
	}
}
