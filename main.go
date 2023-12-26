package main

import (
	"fmt"

	"github.com/rwv/rfask-golang/device/receiver"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	rpio.Open()
	defer rpio.Close()

	rx := receiver.New(PIN_ASK_RX, PIN_ASK_EN, MAX_WAVE_LEN, SAMPLE_PERIOD, 3*1e6, 10*1e6)
	count := 0
	for {
		count += 1
		if wave := rx.Receive(); wave != nil {
			fmt.Printf("%v\n", wave)
		}

		fmt.Println(count)
	}
}
