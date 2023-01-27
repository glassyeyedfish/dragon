package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const IMAGE_SIZE = 512
const OFFSET = 0
const START_X, START_Y = IMAGE_SIZE/2 + OFFSET, IMAGE_SIZE/2 - OFFSET
const STEP = 2
const MAX_ITERATIONS = 14

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

func DrawLine(img *image.NRGBA, x1 int, y1 int, x2 int, y2 int, color color.NRGBA) {
	if x1 == x2 {
		if y1 > y2 {
			y1, y2 = y2, y1
		}

		for y := y1; y <= y2; y++ {
			img.Set(x1, y, color)
		}
	} else {
		if x1 > x2 {
			x1, x2 = x2, x1
		}

		for x := x1; x <= x2; x++ {
			img.Set(x, y1, color)
		}
	}
}

func DrawDragon(img *image.NRGBA, instructions []int, color color.NRGBA) {
	cursor_x, cursor_y := START_X, START_Y
	index := 0

	for iter := 1; iter <= MAX_ITERATIONS; iter++ {

		// Follow the instructions
		for ; index < len(instructions); index++ {
			switch instructions[index] {
			case UP:
				DrawLine(img, cursor_x, cursor_y, cursor_x, cursor_y-STEP, color)
				cursor_y -= STEP
			case DOWN:
				DrawLine(img, cursor_x, cursor_y, cursor_x, cursor_y+STEP, color)
				cursor_y += STEP
			case LEFT:
				DrawLine(img, cursor_x, cursor_y, cursor_x-STEP, cursor_y, color)
				cursor_x -= STEP
			case RIGHT:
				DrawLine(img, cursor_x, cursor_y, cursor_x+STEP, cursor_y, color)
				cursor_x += STEP
			}
		}

		// Work backwards to generate next set of instructions
		for i := index - 1; i >= 0; i-- {
			n := (instructions[i] + 1) % 4
			instructions = append(instructions, n)
		}
	}
}

func main() {

	// internal constants
	color_1 := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	color_2 := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	color_3 := color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	color_4 := color.NRGBA{R: 0, G: 0, B: 255, A: 255}

	// Create image full of black pixels
	img := image.NewNRGBA(image.Rect(0, 0, IMAGE_SIZE, IMAGE_SIZE))
	Clear(img)

	// UP -> RIGHT -> DOWN -> LEFT -> UP
	DrawDragon(img, []int{UP}, color_1)
	DrawDragon(img, []int{DOWN}, color_2)
	DrawDragon(img, []int{LEFT}, color_3)
	DrawDragon(img, []int{RIGHT}, color_4)

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
