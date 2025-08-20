# Badger 2040W TinyGo Image Demo

This project demonstrates how to use [TinyGo](https://tinygo.org/) to
perform graphical operations such as drawing a raw byte array image onto the [Pimoroni Badger
2040W](https://shop.pimoroni.com/products/badger-2040) e-ink display.

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
    ├── button     # Button Manager (polling-based buttons, no interrupts)
    ├── cmd        # Programs such as blink, single_img, slideshow
    ├── device     # Hardware setup (pins, SPI, display)
    ├── output     # Storage of PNG files and slice byte arrays
    ├── syserror   # Utility to render error messages on screen

## Building and Flashing

To build and flash directly to the Badger 2040W:

``` bash
tinygo flash -target=badger2040-w .
```

## Modifying Images

Bitmaps stored in /output defined in **`filename.go`** (filename is a placeholder, use whatever name you prefer):

``` go

var FileName = []byte{ ... }
```

You can replace this byte array with your own generated image data.\
This project uses `pixel.NewImageFromBytes[pixel.Monochrome]` to
interpret the slice.

I wrote a separate Go program to help with the monochrome PNG to byte array conversion. You can find it [here](https://github.com/eithansmith/image2bytes).  

## Usage

1.  Flash the firmware to your Badger 2040W.\
2.  The device clears the screen, draws the screen (based on cmd program logic), waits 5
    seconds, and then sleeps.\
3.  If an error occurs, it will be displayed as text on the screen
    instead of crashing silently.

## License

MIT.
