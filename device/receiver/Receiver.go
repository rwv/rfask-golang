package receiver

import "github.com/stianeikeland/go-rpio/v4"

type Receiver struct {
	PinRX        uint8
	PinEN        uint8
	MaxLen       int64
	SamplePeriod int64 // Sample period in nanoseconds
	MinGap       int64
	MaxGap       int64
	pin          *rpio.Pin
	SampleTime   int64 // Sample Time in nanosecond
	EdgeTime     int64 // Edge Time in nanosecond
	Bit          uint8
}

func New(pinRX, pinEN uint8, maxLen int64, samplePeriodInMillSecond, minGapMs, maxGapMs float64) *Receiver {
	pin := rpio.Pin(pinRX)
	pin.Input()

	return &Receiver{
		pin:    &pin,
		PinRX:  pinRX,
		PinEN:  pinEN,
		MaxLen: maxLen,
		// SamplePeriod:               samplePeriod,
		SamplePeriod: int64(samplePeriodInMillSecond * 1e6),
		MinGap:       int64(minGapMs * 1e6),
		MaxGap:       int64(maxGapMs * 1e6),
	}
}
