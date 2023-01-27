package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const IMAGE_SIZE = 128
const START_X, START_Y = IMAGE_SIZE / 2, IMAGE_SIZE / 2
const STEP = 5
const MAX_ITERATIONS = 7

const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

func Clear(img *image.NRGBA) {
	for y := 0; y < IMAGE_SIZE; y++ {
		for x := 0; x < IMAGE_SIZE; x++ {
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
	img := image.NewNRGBA(image.Rect(0, 0, IMAGE_SIZE, IMAGE_SIZE))

	cursor_x, cursor_y := START_X, START_Y
	index := 0

	// Fill the image with black pixels
	Clear(img)

	// UP -> RIGHT -> DOWN -> LEFT -> UP
	// Drawing instructions
	instructions := []int{UP}

	for iter := 1; iter <= MAX_ITERATIONS; iter++ {

		// Follow the instructions
		for ; index < len(instructions); index++ {
			switch instructions[index] {
			case UP:
				DrawLine(img, cursor_x, cursor_y, cursor_x, cursor_y-STEP)
				cursor_y -= STEP
			case DOWN:
				DrawLine(img, cursor_x, cursor_y, cursor_x, cursor_y+STEP)
				cursor_y += STEP
			case LEFT:
				DrawLine(img, cursor_x, cursor_y, cursor_x-STEP, cursor_y)
				cursor_x -= STEP
			case RIGHT:
				DrawLine(img, cursor_x, cursor_y, cursor_x+STEP, cursor_y)
				cursor_x += STEP
			}
		}

		// Work backwards to generate next set of instructions
		for i := index - 1; i >= 0; i-- {
			n := (instructions[i] + 1) % 4
			instructions = append(instructions, n)
		}
	}

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
