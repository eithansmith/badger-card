package main

import (
	"machine"
	"time"
)

// main (blink) performs a simple blink test of the LED on the badger 2040w
func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
