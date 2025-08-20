package syserror

import (
	"badger_card/device"
	"image/color"
	"os"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

func Write(dev *device.Device, errIn error) {
	// Clear screen
	dev.ClearBuffer()

	lines := splitLines(errIn.Error(), 20)

	cwd, _ := os.Getwd()
	lines = append(lines, cwd)

	// Print each line with vertical spacing
	startY := 20
	lineHeight := 16
	for i, line := range lines {
		y := startY + i*lineHeight
		tinyfont.WriteLine(dev, &freemono.Bold9pt7b, 10, int16(y), line, color.RGBA{R: 1, G: 1, B: 1, A: 255}) // black text
	}

	_ = dev.Display()
}

func splitLines(s string, maxLen int) []string {
	var lines []string
	for i := 0; i < len(s); i += maxLen {
		end := i + maxLen
		if end > len(s) {
			end = len(s)
		}
		lines = append(lines, s[i:end])
	}
	return lines
}
