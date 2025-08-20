package main

import (
	"badger_card/device"
	"badger_card/output"
	"badger_card/output/single_img"
	"badger_card/syserror"
	"log"
	"time"

	"tinygo.org/x/drivers/pixel"
)

func drawScreen(dev *device.Device) error {
	var err error

	defer func() {
		if err != nil {
			syserror.Write(dev, err)
		}
	}()

	image := pixel.NewImageFromBytes[pixel.Monochrome](output.Width, output.Height, single_img.Simple)

	err = dev.DrawBitmap(0, 0, image)
	if err != nil {
		return err
	}

	return nil
}

// main (single image) draws a single image as stored in output/simple
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

	// Init screen
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
