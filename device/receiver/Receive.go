package receiver

import (
	"time"

	"github.com/rwv/rfask-golang/wave"
	"github.com/stianeikeland/go-rpio/v4"
)

func (receiver *Receiver) Receive() *wave.BitWave {
	wave := wave.New(make([]int64, 0), 1)
	ts := wave.TimeStamp
	b := receiver.pin.Read()
	now := time.Now().UnixNano()

	if ch := waitForEdge(receiver.pin, rpio.AnyEdge, receiver.MaxGap); !ch {
		return nil
	}

	b0 := b
	t0 := now
	ts = append(ts, t0)

	if ch := waitForEdge(receiver.pin, rpio.AnyEdge, receiver.MaxGap); !ch {
		return nil
	}

	b = receiver.pin.Read()
	now = time.Now().UnixNano()
	if b == b0 {
		return nil
	}

	wave.StartBit = uint8(b0)
	receiver.Bit = uint8(b)
	receiver.SampleTime = now
	receiver.EdgeTime = now
	ts = append(ts, now)

	// while true
	for {
		b, now := receiver.GetSample()
		if b == receiver.Bit {
			if now-receiver.EdgeTime > receiver.MinGap {
				if b == 0 {
					ts = append(ts, now)
				}
				if len(ts) > 5 {
					return wave
				} else {
					return nil
				}
			}
		} else {
			if (now - receiver.EdgeTime) < receiver.SamplePeriod {
				return nil
			}
			receiver.Bit = b
			receiver.EdgeTime = now
			ts = append(ts, now)
			if len(ts) > int(receiver.MaxLen) {
				return nil
			}
		}
	}

}

func waitForEdge(pin *rpio.Pin, edge rpio.Edge, timeout int64) bool {
	pin.Detect(edge)
	timeoutNow := time.Now().UnixNano()
	for {
		detected := pin.EdgeDetected()
		if detected {
			return true
		}
		if time.Now().UnixNano()-timeoutNow > timeout {
			return false
		}
	}
}
