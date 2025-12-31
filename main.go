package main

import (
	"fmt"
	"os"

	"github.com/jfreymuth/pulse"
)

const baseRate = 44100

func main() {
	c, err := pulse.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	stream, err := c.NewPlayback(pulse.Float32Reader(synth), pulse.PlaybackLatency(0.1))
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

var coreTick uint64

var notes = []Tone{
	E6, F5,
}

var instrument = load("WaveForms/test.png")
var seconds int

var tape = map[uint64][]Note{
	0: {
		{Tone: C4, End: 30000},
	},
	60000: {
		{Tone: E4, End: 20000},
	},
}
var nextNoteID uint64

var activeNotes = map[uint64]Note{}

func synth(out []float32) (int, error) {

	//fmt.Println("synth called:", len(out))
	for i := range out {
		out[i] = 0.0
		note, ok := tape[coreTick]
		if ok {
			delete(tape, coreTick)
			for _, note := range note {
				note.id = nextNoteID
				note.Start = coreTick
				note.End += coreTick
				nextNoteID++
				activeNotes[note.id] = note
				fmt.Printf("Note %s starting at %d and going till %d", note.Tone.Name, coreTick, note.End)
			}
		}
		for id, n := range activeNotes {
			x := instrument.Value(coreTick, n, nil)
			if n.End < coreTick {
				delete(activeNotes, id)
				fmt.Printf("Note %s ending at %d. %d left\n", n.Tone.Name, coreTick, len(activeNotes))
			}
			out[i] += x * 0.1
		}
		coreTick++
		if len(tape) == 0 && len(activeNotes) == 0 {
			fmt.Println("Done")
			os.Exit(0)
		}
	}
	return len(out), nil
}

type Note struct {
	id      uint64
	Tone    Tone
	Start   uint64
	End     uint64
	Attack  uint64
	Decay   uint64
	Sustain float32
	Release uint64
}
