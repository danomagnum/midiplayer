package main

import (
	"fmt"

	"github.com/jfreymuth/pulse"
)

func main() {
	c, err := pulse.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	stream, err := c.NewPlayback(pulse.Float32Reader(synth), pulse.PlaybackLatency(.1))
	if err != nil {
		fmt.Println(err)
		return
	}

	stream.Start()
	stream.Drain()
	defer stream.Close()
	for {
		//fmt.Println("Underflow:", stream.Underflow())
		if stream.Error() != nil {
			fmt.Println("Error:", stream.Error())
			return
		}
	}
}

var phase float64
var tick int

var notes = []Note{
	E6, F5,
}

var instrument = load("WaveForms/test.png")
var seconds int

func synth(out []float32) (int, error) {

	//fmt.Println("synth called:", len(out))
	for i := range out {
		n := notes[int(tick%len(notes))]
		x := instrument.Value(phase, n, nil)
		out[i] = x * 0.1
		phase += 1.0
		if phase >= 44100 {
			phase = 0
			seconds++
		}
		if seconds == 1 {
			fmt.Println("Note change:", n.Frequency, n.Name)
			tick++
			seconds = 0
		}
	}
	return len(out), nil
}
