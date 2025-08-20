package main

import (
	"badger_card/device"
	"image/color"
	"log"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

// main (text) displays a simple line of text
func main() {
	var err error

	dev, err := device.New()
	if err != nil {
		log.Fatal(err)
	}

	dev.ClearBuffer()

	text := "THERE IS NO SPOON"

	tinyfont.WriteLine(dev, &freemono.Bold12pt7b, 25, 60, text, color.RGBA{R: 1, G: 1, B: 1, A: 255})

	_ = dev.Display()

	_ = dev.Sleep(false)
}
