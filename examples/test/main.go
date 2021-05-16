package main

import (
	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/color"
)

func main() {
	// create a new frame buffer
	fb := screen.NewFrameBuffer()

	// turn on LEDs on the first row
	fb.SetPixel(0, 0, color.Red)
	fb.SetPixel(1, 1, color.Blue)

	screen.Draw(fb)
}
