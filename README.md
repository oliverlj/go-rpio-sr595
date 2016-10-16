go-rpio-sr595
=======

Extension of [go-rpio](https://raw.githubusercontent.com/stianeikeland/go-rpio) project for your Pi!
This project bring the support of the shift register 74HC595.

## Releases ##
- In progress

## Usage by example ##

```go 
go get "github.com/olivierlj/go-rpio-sr595"
```

Example :
```go
var (
        numberRegisterPins = 8  // Number of register pins
        dataPin = rpio.Pin(22)
        clockPin = rpio.Pin(17)
        latchPin = rpio.Pin(27)
)

func main() {
        // Open and map memory to access gpio, check for errors
        if err := rpio.Open(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        // Unmap gpio memory when done
        defer rpio.Close()

        sr595.Setup(clockPin, dataPin, latchPin, numberRegisterPins) // Setup the shift register with pins and number of pins
        sr595.Reset() // Set all pins to Low, with immediate effect
        sr595.SetRegisterPin(1, rpio.High) // Set first pin (pin are 1-index) to High
        sr595.WriteRegisters() // Apply value to shift registers
}
```
