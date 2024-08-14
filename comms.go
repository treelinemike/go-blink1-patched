package blink1

/*
	#cgo LDFLAGS: -lusb
	#include <usb.h>
*/
import "C"

import (
	"time"

	"github.com/treelinemike/go-blink1-patched/libusb"
)

func fadeToRgbBlink1(device *Device, fadeTime time.Duration, red, green, blue, led uint8) (bytesWritten int) {
	bytesWritten = libusb.SendBlink1Command(device.Device, toMs(fadeTime), red, blue, green, led)
	return
}
