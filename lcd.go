// Package wipi a wrapper for the wiringPi c library on Raspberry Pi
package wipi

// #cgo CFLAGS: -IWiringPi/wiringPi
// #cgo LDFLAGS: -LWiringPi/wiringPi -lwiringPi
// #include "stdint.h"
// #include "lcd.h"
import "C" 

import "fmt"

type LCD struct {
	fd,
	rows,
	cols int
}

func (lcd *LCD) Clear() {
	C.lcdClear(C.int(lcd.fd))
}

func (lcd *LCD) Home() {
	lcd.SetCursor(1, 1)
}

func (lcd *LCD) SetCursor(x, y int) {
	C.lcdPosition(C.int(lcd.fd), C.int(x), C.int(y))
}

func (lcd *LCD) Print(message string) {
	C.lcdPuts(C.int(lcd.fd), C.CString(message))
}

func NewLCD(rs, enable, d4, d5, d6, d7 int) (lcd *LCD, err error) {
	err = nil
	lcd = new(LCD)
	fmt.Println("In NewLCD")
	var fd int = int(C.lcdInit(
		2, 16, 4, C.int(rs), C.int(enable),
		C.int(d4), C.int(d5), C.int(d6), C.int(d7), C.int(0), C.int(0), 
		C.int(0), C.int(0)))
	lcd.fd = fd
	fmt.Println("FD: %d", fd)
	return lcd, err
}
