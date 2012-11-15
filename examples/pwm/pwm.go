package main

import ( 
	"fmt"
	"github.com/steve918/wipi"
	"time"
)

const (
	PWM_LED = 7
)

func main() {
	r := wipi.WiringPiSetupGpio()
	if r == 1 {
		fmt.Println("wiringPiSetupGpio complete")
	}
	wipi.PinMode(PWM_LED, wipi.PWM_OUTPUT)
	for {
		for i :=0; i < 1024; i++ {
			wipi.PwmWrite(PWM_LED, i)
			time.Sleep(10 * time.Millisecond)
		}

		for i := 1023; i >= 0; i-- {
			wipi.PwmWrite(PWM_LED, i)
			time.Sleep(10 * time.Millisecond)
		}

	}
}
