package system

import (
	"strconv"

	"github.com/szampardi/gobot"
)

// sysfsDitalPinHandler represents the sysfs implementation
type sysfsDigitalPinAccess struct {
	sfa *sysfsFileAccess
}

// gpiodDigitalPinAccess represents the character device implementation
type gpiodDigitalPinAccess struct {
	fs    filesystem
	chips []string
}

func (h *sysfsDigitalPinAccess) isSupported() bool {
	// currently this is supported by all Kernels
	return true
}

func (h *sysfsDigitalPinAccess) createPin(chip string, pin int,
	o ...func(gobot.DigitalPinOptioner) bool,
) gobot.DigitalPinner {
	return newDigitalPinSysfs(h.sfa, strconv.Itoa(pin), o...)
}

func (h *sysfsDigitalPinAccess) setFs(fs filesystem) {
	h.sfa = &sysfsFileAccess{fs: fs, readBufLen: 2}
}

func (h *gpiodDigitalPinAccess) isSupported() bool {
	chips, err := h.fs.find("/dev", "gpiochip")
	if err != nil || len(chips) == 0 {
		return false
	}
	h.chips = chips
	return true
}

func (h *gpiodDigitalPinAccess) createPin(chip string, pin int,
	o ...func(gobot.DigitalPinOptioner) bool,
) gobot.DigitalPinner {
	return newDigitalPinGpiod(chip, pin, o...)
}

func (h *gpiodDigitalPinAccess) setFs(fs filesystem) {
	h.fs = fs
}
