package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const w, h = 256, 256
const hw, hh = w / 2, h / 2
const step = 5

const (
	UP    = iota
	DOWN  = iota
	RIGHT = iota
	LEFT  = iota
)

func Clear(img *image.NRGBA) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.NRGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 255,
			})
		}
	}
}

func DrawLine(img *image.NRGBA, x1 int, y1 int, x2 int, y2 int) {
	if x1 == x2 {
		if y1 > y2 {
			y1, y2 = y2, y1
		}

		for y := y1; y <= y2; y++ {
			img.Set(x1, y, color.NRGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			})
		}
	} else {
		if x1 > x2 {
			x1, x2 = x2, x1
		}

		for x := x1; x <= x2; x++ {
			img.Set(x, y1, color.NRGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			})
		}
	}
}

func main() {

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, w, h))

	// Set cursor to the middle of the image
	cx, cy := hw, hh

	// Fill the image with black pixels
	Clear(img)

	DrawLine(img, cx, cy, cx, cy+step)
	cy = cy + step
	DrawLine(img, cx, cy, cx+step, cy)

	// Save it to the image file
	f, err := os.Create("dragon.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
