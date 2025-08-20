# Badger 2040W TinyGo Image Demo

This project demonstrates how to use [TinyGo](https://tinygo.org/) to
draw a raw byte array image onto the [Pimoroni Badger
2040W](https://shop.pimoroni.com/products/badger-2040) e-ink display.\
It sets up the device, loads a `[]byte` image, and renders it
fullscreen.

## Features

-   Initializes the UC8151 e-ink controller over SPI.
-   Configures LED and button pins (A, B, C, Up, Down).
-   Draws an embedded bitmap at 296×128 pixels or a slideshow of pages, controllable by up/down button press.
-   Shows errors directly on the e-ink display (with TinyFont).
-   Puts the device to sleep after drawing, preserving battery life.

## Prerequisites

-   [TinyGo](https://tinygo.org/) (≥0.28 recommended).
-   A Pimoroni **Badger 2040W** or compatible RP2040 board with UC8151
    display.
-   USB cable and flashing tool (`tinygo flash` or
    [picotool](https://github.com/raspberrypi/picotool)).

## Project Structure

    .
    ├── button     # Button Manager (polling-based buttons (no interrupts)
    ├── cmd        # Programs such as blink, single_img, slideshow
    ├── device     # Hardware setup (pins, SPI, display)
    ├── output     # Storage of PNG files and slice byte arrays
    ├── syserror   # Utility to render error messages on screen

## Building and Flashing

To build and flash directly to the Badger 2040W:

``` bash
tinygo flash -target=badger2040-w .
```

## Modifying the Image

The bitmap stored in /output is defined in **`filename.go`**:

``` go

var FileName = []byte{ ... }
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
