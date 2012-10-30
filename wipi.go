// Package wipi a wrapper for the wiringPi c library on Raspberry Pi
package wipi

// #cgo LDFLAGS: -lwiringPi
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

//extern int  wiringPiSetup       (void) ;
func WiringPiSetup() int {
	var r C.int = C.wiringPiSetup()
	return int(r)
}


//extern int  wiringPiSetupSys    (void) ;
func WiringPiSetupSys() int {
	var r C.int = C.wiringPiSetupSys()
	return int(r)
}


//extern int  wiringPiSetupGpio   (void) ;
func WiringPiSetupGpio() int {
	var r C.int = C.wiringPiSetupGpio()
	return int(r)
}


//extern int  piBoardRev          (void) ;
func PiBoardRev() int {
	var r C.int = C.piBoardRef()
	return int(r)
}

//extern int  wpiPinToGpio        (int wpiPin) ;
func WpiPinToGpio() int {
	var r C.int = C.wpiPinToGpio()
	return int(r)
}

//extern void (*pinMode)           (int pin, int mode) ;
func PinMode(pin, mode int) {
	C.pinmode(C.int(pin), C.int(mode))
} 

//extern void (*pullUpDnControl)   (int pin, int pud) ;

//extern void (*digitalWrite)      (int pin, int value) ;
func DigitalWrite(pin, value int) {
	C.digitalWrite(C.int(pin), C.int(value));
}

//extern void (*pwmWrite)          (int pin, int value) ;
func PwmWrite(pin, value int) {
	C.pwmlWrite(C.int(pin), C.int(value));
}

//extern int  (*digitalRead)       (int pin) ;
func DigitalRead(pin int) int {
	var r C.int = C.digitalRead(C.int(pin))
	return int(r)
}

//extern void (*pwmSetMode)        (int mode) ;
func PwmSetMode(mode int) {
	C.pwmSetMode(C.int(mode))
}

//extern void (*pwmSetRange)       (unsigned int range) ;
func PwmSetRange(range uint) {
	C.pwmSetRange(C.uint(range))
}

//extern void (*pwmSetClock)       (int divisor) ;
func PwmSetClock(divisor int) {
	C.pwmSetClock(C.int(divisor))
}


