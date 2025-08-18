package device

import (
	"fmt"
	"machine"
	"tinygo.org/x/drivers/uc8151"
)

type Device struct {
	ActLED  machine.Pin
	ABtn    machine.Pin
	BBtn    machine.Pin
	CBtn    machine.Pin
	UpBtn   machine.Pin
	DownBtn machine.Pin

	*uc8151.Device
}

func New() (*Device, error) {

	machine.ENABLE_3V3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.ENABLE_3V3.High()

	actLED := machine.LED
	actLED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	aBtn := machine.BUTTON_A
	aBtn.Configure(machine.PinConfig{Mode: machine.PinInput})
	bBtn := machine.BUTTON_B
	bBtn.Configure(machine.PinConfig{Mode: machine.PinInput})
	cBtn := machine.BUTTON_C
	cBtn.Configure(machine.PinConfig{Mode: machine.PinInput})

	upBtn := machine.BUTTON_UP
	upBtn.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	err := upBtn.SetInterrupt(machine.PinRising, func(p machine.Pin) {
		actLED.Set(true)
	})
	if err != nil {
		return nil, fmt.Errorf("setting up interrupt: %w", err)
	}
	downBtn := machine.BUTTON_DOWN
	downBtn.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	err = downBtn.SetInterrupt(machine.PinRising, func(p machine.Pin) {
		actLED.Set(false)
	})
	if err != nil {
		return nil, fmt.Errorf("setting up interrupt: %w", err)
	}

	display, err := createDisplay()
	if err != nil {
		return nil, fmt.Errorf("initializing display: %w", err)
	}

	return &Device{
		ABtn:    aBtn,
		BBtn:    bBtn,
		CBtn:    cBtn,
		UpBtn:   upBtn,
		DownBtn: downBtn,

		Device: display,
		ActLED: actLED,
	}, nil
}

func createDisplay() (*uc8151.Device, error) {
	err := machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12_000_000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})
	if err != nil {
		return nil, fmt.Errorf("initializing spi0: %w", err)
	}

	display := uc8151.New(
		machine.SPI0,
		machine.EPD_CS_PIN,
		machine.EPD_DC_PIN,
		machine.EPD_RESET_PIN,
		machine.EPD_BUSY_PIN,
	)
	display.Configure(uc8151.Config{
		Rotation: uc8151.ROTATION_270,
		Speed:    uc8151.MEDIUM,
		Blocking: true,
	})

	return &display, nil
}
