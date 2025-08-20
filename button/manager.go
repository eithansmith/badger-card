package button

import (
	"machine"
	"time"
)

type Manager struct {
	btns      []btnDef
	onShort   func(name string)
	onLong    func(name string)
	debounce  time.Duration
	longPress time.Duration
	period    time.Duration
}

type btnDef struct {
	pin      machine.Pin
	name     string
	lastEdge time.Time
	downAt   time.Time
	isDown   bool
}

func NewManager(onShort, onLong func(string)) *Manager {
	return &Manager{
		onShort:   onShort,
		onLong:    onLong,
		debounce:  40 * time.Millisecond,
		longPress: 600 * time.Millisecond,
		period:    8 * time.Millisecond, // ~125 Hz
		btns: []btnDef{
			{pin: machine.BUTTON_A, name: "A"},
			{pin: machine.BUTTON_B, name: "B"},
			{pin: machine.BUTTON_C, name: "C"},
			{pin: machine.BUTTON_UP, name: "UP"},
			{pin: machine.BUTTON_DOWN, name: "DOWN"},
		},
	}
}

// StartPolling buttons (no interrupts). A/B/C: PinInput; UP/DOWN: PinInputPulldown.
// If any button acts inverted on your board, flip its Mode to PinInputPullup.
func (m *Manager) StartPolling() {
	for i := range m.btns {
		b := &m.btns[i]
		mode := machine.PinInput
		if b.name == "UP" || b.name == "DOWN" {
			mode = machine.PinInputPulldown
		}
		b.pin.Configure(machine.PinConfig{Mode: mode})
	}

	go func() {
		t := time.NewTicker(m.period)
		defer t.Stop()

		for range t.C {
			now := time.Now()
			for i := range m.btns {
				b := &m.btns[i]
				level := b.pin.Get() // HIGH means pressed for pulldown-wired

				// Rising edge (press)
				if level && !b.isDown {
					if now.Sub(b.lastEdge) >= m.debounce {
						b.isDown = true
						b.downAt = now
						b.lastEdge = now
					}
					continue
				}

				// Falling edge (release)
				if !level && b.isDown {
					if now.Sub(b.lastEdge) >= m.debounce {
						b.isDown = false
						held := now.Sub(b.downAt)
						b.lastEdge = now

						if held >= m.longPress {
							if m.onLong != nil {
								m.onLong(b.name)
							}
						} else if m.onShort != nil {
							m.onShort(b.name)
						}
					}
				}
			}
		}
	}()
}
