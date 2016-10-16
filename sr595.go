package sr595

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

var (
	numOfRegisterPins int
	dataPin           = rpio.Pin(22)
	clockPin          = rpio.Pin(17)
	latchPin          = rpio.Pin(27)
	pins              []rpio.State
)

func Setup(clockPin, dataPin, latchPin rpio.Pin, numberOfPins int) {
	dataPin.Low()
	clockPin.Low()
	latchPin.High()
	dataPin.Output()
	clockPin.Output()
	latchPin.Output()
	numOfRegisterPins = numberOfPins
}

func Reset() {
	clearRegisters()
	WriteRegisters()
}

func clearRegisters() {
	for i := 0; i < numOfRegisterPins; i++ {
		pins[i] = rpio.Low
	}
}

func WriteRegisters() {
	latchPin.Low()
	time.Sleep(time.Microsecond)
	for i := numOfRegisterPins - 1; i >= 0; i-- {
		dataPin.Write(pins[i])
		clockPin.High()
		time.Sleep(time.Microsecond)
		clockPin.Low()
		time.Sleep(time.Microsecond)
	}
	latchPin.High()
	time.Sleep(time.Microsecond)
}

//set an individual pin HIGH or LOW
func SetRegisterPin(index int, value rpio.State) {
	pins[index-1] = value
}
