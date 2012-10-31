// Package wipi a wrapper for the wiringPi c library on Raspberry Pi
package wipi

// #cgo CFLAGS: -IWiringPi/wiringPi
// #cgo LDFLAGS: -LWiringPi/wiringPi -lwiringPi
// #include "wiringPi.h"
import "C"


const (
	HIGH = 1
	LOW = 0
	MSBFIRST = 1
	LSBFIRST = 0
	WPI_MODE_PINS = 0
	WPI_MODE_GPIO = 1
	WPI_MODE_SYS = 2
	MODE_PINS = 0
	MODE_GPIO = 1
	MODE_SYS = 2
	INPUT = 0
	OUTPUT = 1
	PWM_OUTPUT = 2
	PUD_OFF = 0
	PUD_DOWN = 1
	PUD_UP = 2
	MODE = 0
)

func WiringPiSetup() int {
	var r C.int = C.wiringPiSetup()
	return int(r)
}

func WiringPiSetupSys() int {
	var r C.int = C.wiringPiSetupSys()
	return int(r)
}

func WiringPiSetupGpio() int {
	var r C.int = C.wiringPiSetupGpio()
	return int(r)
}

func PiBoardRev() int {
	var r C.int = C.piBoardRev()
	return int(r)
}

func WpiPinToGpio(pin int) int {
	var r C.int = C.wpiPinToGpio(C.int(pin))
	return int(r)
}

func PinMode(pin, mode int) {
	C._pinMode(C.int(pin), C.int(mode))
} 

func PullUpDnControl(pin, pud int) {
	C._pullUpDnControl(C.int(pin), C.int(pud));
}

func DigitalWrite(pin, value int) {
	C._digitalWrite(C.int(pin), C.int(value));
}

func PwmWrite(pin, value int) {
	C._pwmWrite(C.int(pin), C.int(value));
}

func DigitalRead(pin int) int {
	var r C.int = C._digitalRead(C.int(pin))
	return int(r)
}

func PwmSetMode(mode int) {
	C._pwmSetMode(C.int(mode))
}

func PwmSetRange(pwmrange uint) {
	C._pwmSetRange(C.uint(pwmrange))
}

func PwmSetClock(divisor int) {
	C._pwmSetClock(C.int(divisor))
}
