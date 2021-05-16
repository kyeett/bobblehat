package main

import (
	"fmt"
	"github.com/kyeett/bobblehat/sense/screen"
	"github.com/kyeett/bobblehat/sense/screen/color"
	"time"
)

func main() {
	fmt.Println("starting")
	// create a new frame buffer
	fb := screen.NewFrameBuffer()

	// turn on LEDs on the first row
	fb.SetPixel(0, 0, color.Red)
	fb.SetPixel(1, 1, color.Red)

	fmt.Println("red")
	for i := uint16(0); i < 256; i++ {
		draw(uint8(i), 0, 0)
		time.Sleep(100 * time.Millisecond)
	}

	//f, err := os.Create(backBuffer)
	//if err != nil {
	//	return err
	//}
	//defer f.Close()
	//
	//return binary.Write(f, binary.LittleEndian, fb.Texture.Pixels)
}

func draw(red, green, blue uint8) {
	r := red >> 3
	screen.DrawAny(uint16(r) << 11)
}
