package receiver

import (
	"time"
)

func (receiver *Receiver) GetSample() (uint8, int64) {
	receiver.SampleTime += receiver.SamplePeriod
	now := time.Now().UnixNano()
	wait := receiver.SampleTime - now
	if wait > 0 {
		time.Sleep(time.Duration(wait))
	}

	res := receiver.pin.Read()
	now = time.Now().UnixNano()

	return uint8(res), now

	// Write your code here
}
