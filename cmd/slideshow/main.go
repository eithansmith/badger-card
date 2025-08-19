package main

import (
	"badger_card/button"
	"badger_card/device"
	"badger_card/output"
	"badger_card/output/slideshow"
	"log"
	"time"

	"tinygo.org/x/drivers/pixel"
)

// Globals kept small and simple for callbacks.
var (
	dev    *device.Device
	pages  = slideshow.Pages
	pageIx = 0
)

// Render current page
func showPage() error {
	img := pixel.NewImageFromBytes[pixel.Monochrome](output.Width, output.Height, pages[pageIx])
	if err := dev.DrawBitmap(0, 0, img); err != nil {
		return err
	}
	return dev.Display()
}

// Clear -> Display (white) -> ShowPage : good for ghosting cleanup
func clearThenShow() error {
	dev.ClearBuffer()
	if err := dev.Display(); err != nil {
		return err
	}
	return showPage()
}

func nextPage() {
	if len(pages) < 2 {
		return // avoid re-drawing same page (prevents darkening)
	}
	old := pageIx
	pageIx = (pageIx + 1) % len(pages)
	if pageIx != old {
		_ = showPage()
	}
}

func prevPage() {
	if len(pages) < 2 {
		return
	}
	old := pageIx
	pageIx = (pageIx - 1 + len(pages)) % len(pages)
	if pageIx != old {
		_ = showPage()
	}
}

func onShortPress(name string) {
	switch name {
	case "UP":
		nextPage()
	case "DOWN":
		prevPage()
	case "A":
		// LED blink as feedback
		dev.ActLED.High()
		time.Sleep(60 * time.Millisecond)
		dev.ActLED.Low()
	case "B":
		// Strong refresh to fight ghosting
		_ = clearThenShow()
	case "C":
		// Quick clear
		dev.ClearBuffer()
		_ = dev.Display()
	}
}

func onLongPress(name string) {
	switch name {
	case "C":
		_ = dev.Sleep(false) // sleep on long-press C
	}
}

// main (slideshow) loads a slice of images in output/slideshow and toggles them via up and down buttons
func main() {
	var err error
	dev, err = device.New()
	if err != nil {
		log.Fatal(err)
	}

	// Initial paint
	dev.ClearBuffer()
	err = dev.Display()
	if err != nil {
		log.Fatal(err)
	}

	err = showPage()
	if err != nil {
		log.Fatal(err)
	}

	// Start polled buttons
	bm := button.NewManager(onShortPress, onLongPress)
	bm.StartPolling()

	// Sit idle; callbacks do the work
	for {
		time.Sleep(500 * time.Millisecond)
	}
}
