package main

import ( 
	"fmt"
	"github.com/steve918/wipi"
	"time"
	"strconv"
)

const (
	BLINK_LED = 7
)

func main() {
	version := wipi.PiBoardRev()
	fmt.Println("Rpi board version: " + strconv.Itoa(version))
	r := wipi.WiringPiSetupGpio()
	if r == 1 {
		fmt.Println("wiringPiSetupGpio complete")
	}
	wipi.PinMode(BLINK_LED, wipi.OUTPUT)
	for {
		wipi.DigitalWrite(BLINK_LED, wipi.HIGH)
		time.Sleep(100 * time.Millisecond)
		wipi.DigitalWrite(BLINK_LED, wipi.LOW)
		time.Sleep(100 * time.Millisecond)
	}
}
