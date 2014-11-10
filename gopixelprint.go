package main

import (
	"fmt"

	"github.com/PhilmacFLy/gopixelprint/gohpglpixel"

	"log"
)

func main() {
	p := &gohpglpixel.Pixelart{}
	p.SetDim(16, 16)

	p.SetPixel(5, 0, 1)
	p.SetPixel(6, 0, 1)
	p.SetPixel(7, 0, 1)
	p.SetPixel(8, 0, 1)
	p.SetPixel(9, 0, 1)
	p.SetPixel(10, 0, 1)

	p.SetPixel(3, 1, 1)
	p.SetPixel(4, 1, 1)
	p.SetPixel(5, 1, 4)
	p.SetPixel(6, 1, 4)
	p.SetPixel(7, 1, 4)
	p.SetPixel(8, 1, 4)
	p.SetPixel(11, 1, 1)
	p.SetPixel(12, 1, 1)

	p.SetPixel(2, 2, 1)
	p.SetPixel(5, 2, 4)
	p.SetPixel(6, 2, 4)
	p.SetPixel(7, 2, 4)
	p.SetPixel(8, 2, 4)
	p.SetPixel(13, 2, 1)

	p.SetPixel(1, 3, 1)
	p.SetPixel(4, 3, 4)
	p.SetPixel(5, 3, 4)
	p.SetPixel(6, 3, 4)
	p.SetPixel(7, 3, 4)
	p.SetPixel(8, 3, 4)
	p.SetPixel(9, 3, 4)
	p.SetPixel(14, 3, 1)

	p.SetPixel(1, 4, 1)
	p.SetPixel(3, 4, 4)
	p.SetPixel(4, 4, 4)
	p.SetPixel(9, 4, 4)
	p.SetPixel(10, 4, 4)
	p.SetPixel(14, 4, 1)

	p.SetPixel(0, 5, 1)
	p.SetPixel(1, 5, 4)
	p.SetPixel(2, 5, 4)
	p.SetPixel(3, 5, 4)
	p.SetPixel(10, 5, 4)
	p.SetPixel(11, 5, 4)
	p.SetPixel(12, 5, 4)
	p.SetPixel(13, 5, 4)
	p.SetPixel(14, 5, 4)
	p.SetPixel(15, 5, 1)

	p.SetPixel(0, 6, 1)
	p.SetPixel(1, 6, 4)
	p.SetPixel(2, 6, 4)
	p.SetPixel(3, 6, 4)
	p.SetPixel(10, 6, 4)
	p.SetPixel(11, 6, 4)
	p.SetPixel(14, 6, 4)
	p.SetPixel(15, 6, 1)

	p.SetPixel(0, 7, 1)
	p.SetPixel(2, 7, 4)
	p.SetPixel(3, 7, 4)
	p.SetPixel(10, 7, 4)
	p.SetPixel(15, 7, 1)

	p.SetPixel(0, 8, 1)
	p.SetPixel(3, 8, 4)
	p.SetPixel(4, 8, 4)
	p.SetPixel(9, 8, 4)
	p.SetPixel(10, 8, 4)
	p.SetPixel(15, 8, 1)

	p.SetPixel(0, 9, 1)
	p.SetPixel(3, 9, 4)
	p.SetPixel(4, 9, 4)
	p.SetPixel(5, 9, 4)
	p.SetPixel(6, 9, 4)
	p.SetPixel(7, 9, 4)
	p.SetPixel(8, 9, 4)
	p.SetPixel(9, 9, 4)
	p.SetPixel(10, 9, 4)
	p.SetPixel(11, 9, 4)
	p.SetPixel(14, 9, 4)
	p.SetPixel(15, 9, 1)

	p.SetPixel(0, 10, 1)
	p.SetPixel(2, 10, 4)
	p.SetPixel(3, 10, 4)
	p.SetPixel(4, 10, 1)
	p.SetPixel(5, 10, 1)
	p.SetPixel(6, 10, 1)
	p.SetPixel(7, 10, 1)
	p.SetPixel(8, 10, 1)
	p.SetPixel(9, 10, 1)
	p.SetPixel(10, 10, 1)
	p.SetPixel(11, 10, 1)
	p.SetPixel(12, 10, 4)
	p.SetPixel(13, 10, 4)
	p.SetPixel(14, 10, 4)
	p.SetPixel(15, 10, 1)

	p.SetPixel(1, 11, 1)
	p.SetPixel(2, 11, 1)
	p.SetPixel(3, 11, 1)
	p.SetPixel(6, 11, 1)
	p.SetPixel(9, 11, 1)
	p.SetPixel(12, 11, 1)
	p.SetPixel(13, 11, 1)
	p.SetPixel(14, 11, 1)

	p.SetPixel(2, 12, 1)
	p.SetPixel(6, 12, 1)
	p.SetPixel(9, 12, 1)
	p.SetPixel(13, 12, 1)

	p.SetPixel(2, 13, 1)
	p.SetPixel(13, 13, 1)

	p.SetPixel(3, 14, 1)
	p.SetPixel(12, 14, 1)

	p.SetPixel(4, 15, 1)
	p.SetPixel(5, 15, 1)
	p.SetPixel(6, 15, 1)
	p.SetPixel(7, 15, 1)
	p.SetPixel(8, 15, 1)
	p.SetPixel(9, 15, 1)
	p.SetPixel(10, 15, 1)
	p.SetPixel(11, 15, 1)

	fmt.Println("asdf")
	p.Print()
	p.SetScaling(10)
	p.SetBorder(false)
	p.SetFilling(4)
	err := p.WritePixelart("schwamerln")
	if err != nil {
		log.Fatal(err)
	}
	p.SetTitle("Schwamerl")
	p.SaveHPGL("blub.hpgl")

	test := &gohpglpixel.Pixelart{}
	test.ReadFile("testfile")
	fmt.Println(test.Canvas)
}
