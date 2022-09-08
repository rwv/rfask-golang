package wave

import "math/rand"

type BitWave struct {
	TimeStamp []int64 // in nano second
	StartBit  uint8
}

//timestamp=None, startbit=1
func New(timestamp []int64, startbit uint8) *BitWave {
	return &BitWave{
		TimeStamp: timestamp,
		StartBit:  startbit,
	}
}

func (bitwave *BitWave) Dither(r int64) {
	timestamp := bitwave.TimeStamp
	// add a dither to the timestamp
	for i := 0; i < len(timestamp); i++ {
		timestamp[i] += int64(rand.Float64() * float64(r))
	}
}

func (bitwave *BitWave) Random(smin, smax, wmin, wmax int64) {
	bitwave.StartBit = uint8(rand.Int63n(2))
	size := rand.Int63n(smax-smin+1) + smin
	t := int64(4e+3)
	// timestamp := [0, t]
	timestamp := make([]int64, 0)
	timestamp = append(timestamp, 0)
	timestamp = append(timestamp, t)
	r := wmax - wmin
	// loop size
	for i := 0; i < int(size); i++ {
		t += int64(rand.Float64()*float64(r)) + wmin
		timestamp = append(timestamp, t)
	}

	t = int64(3e+3)
	timestamp = append(timestamp, t)
	bitwave.TimeStamp = timestamp
}
