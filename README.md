# Badger 2040W TinyGo Image Demo

This project demonstrates how to use [TinyGo](https://tinygo.org/) to
draw a raw byte array image onto the [Pimoroni Badger
2040W](https://shop.pimoroni.com/products/badger-2040) e-ink display.\
It sets up the device, loads a `[]byte` image, and renders it
fullscreen.

## Features

-   Initializes the UC8151 e-ink controller over SPI.
-   Configures LED and button pins (A, B, C, Up, Down).
-   Draws an embedded bitmap (`output.go`) at 296×128 pixels.
-   Shows errors directly on the e-ink display (with TinyFont).
-   Puts the device to sleep after drawing, preserving battery life.

## Prerequisites

-   [TinyGo](https://tinygo.org/) (≥0.28 recommended).
-   A Pimoroni **Badger 2040W** or compatible RP2040 board with UC8151
    display.
-   USB cable and flashing tool (`tinygo flash` or
    [picotool](https://github.com/raspberrypi/picotool)).

## Project Files

    .
    ├── device.go   # Hardware setup (pins, SPI, display)
    ├── error.go    # Utility to render error messages on screen
    ├── main.go     # Entry point: draws image, handles sleep
    ├── output.go   # Embedded []byte image data (296x128 monochrome)

## Building and Flashing

To build and flash directly to the Badger 2040W:

``` bash
tinygo flash -target=pico .
```

Or, to build a UF2 file you can drag-and-drop:

``` bash
tinygo build -o firmware.uf2 -target=pico .
```

(Use `-target=badger2040` if you have a custom board definition
available in your TinyGo installation.)

## Modifying the Image

The bitmap is defined in **`output.go`**:

``` go
const OutputWidth = 296
const OutputHeight = 128

var Output = []byte{ ... }
```

You can replace this byte array with your own generated image data.\
The example uses `pixel.NewImageFromBytes[pixel.Monochrome]` to
interpret the slice.

## Usage

1.  Flash the firmware to your Badger 2040W.\
2.  The device clears the screen, draws the embedded image, waits 5
    seconds, and then sleeps.\
3.  If an error occurs, it will be displayed as text on the screen
    instead of crashing silently.

## License

MIT (or whatever license you prefer).
