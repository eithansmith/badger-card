package main

import (
	"badger_card/device"
	"badger_card/output"
	"badger_card/syserror"
	"image/color"
	"log"
	"time"

	"tinygo.org/x/drivers/pixel"
)

var (
	black       = color.RGBA{R: 1, G: 1, B: 1, A: 255}
	white       = color.RGBA{A: 255}
	w     int16 = 296
	h     int16 = 128
)

func drawScreen(dev *device.Device) error {
	var err error

	defer func() {
		if err != nil {
			syserror.Write(dev, err)
		}
	}()

	image := pixel.NewImageFromBytes[pixel.Monochrome](output.OutputWidth, output.OutputHeight, output.Output)

	err = dev.DrawBitmap(0, 0, image)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var err error

	dev, err := device.New()
	if err != nil {
		log.Fatal(err)
	}

	err = run(dev)
	if err != nil {
		log.Fatal(err)
	}
}

func run(dev *device.Device) error {
	var err error

	dev.ClearBuffer()
	err = dev.Display()
	if err != nil {
		return err
	}

	// Draw screen
	err = drawScreen(dev)
	if err != nil {
		return err
	}

	// Update screen
	err = dev.Display()
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)

	// Sleep to preserve battery
	err = dev.Sleep(false)
	if err != nil {
		return err
	}

	return nil
}
