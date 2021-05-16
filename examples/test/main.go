package main

import (
	"github.com/kyeett/bobblehat/sense/screen"
	"github.com/kyeett/bobblehat/sense/screen/color"
	"time"
)

func main() {
	// create a new frame buffer
	fb := screen.NewFrameBuffer()

	// turn on LEDs on the first row
	fb.SetPixel(0, 0, color.Red)
	fb.SetPixel(1, 1, color.Red)

	for i := uint16(0); i < 65535; i++ {
		//t := screen.FrameBuffer{
		//	Texture: &texture.Texture{
		//		Pixels: []color.Color{color.Color(i), color.Color(i), color.Color(i)},
		//	},
		//}
		//screen.Draw(&t)
		screen.DrawAny(i)
		time.Sleep(1 * time.Millisecond)
	}

	//f, err := os.Create(backBuffer)
	//if err != nil {
	//	return err
	//}
	//defer f.Close()
	//
	//return binary.Write(f, binary.LittleEndian, fb.Texture.Pixels)
}
