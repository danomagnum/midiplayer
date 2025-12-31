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

	for i, t := range Mary {
		start := uint64(i * 10000)
		tape[start] = []Note{
			{Tone: t, End: 10, Attack: 10, Release: 5000, Instrument: &piano},
			{Tone: t, End: 7000, Attack: 3000, Release: 3000, Instrument: &violin},
		}
	}

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

var instrument1 = load("WaveForms/test.png")
var piano = load("WaveForms/piano.png")
var violin = load("WaveForms/violin.png")
var seconds int

var tape = map[uint64][]Note{}
var Mary = []Tone{
	E5, D5, C5, D5, E5, E5, E5,
	D5, D5, D5,
	E5, G5, G5,
	E5, D5, C5, D5, E5, E5, E5, E5,
	D5, D5, E5, D5, C5,
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
				note.ID = nextNoteID
				note.Start = coreTick
				note.End += coreTick
				nextNoteID++
				activeNotes[note.ID] = note
				fmt.Printf("Note %s starting at %d and going till %d", note.Tone.Name, coreTick, note.End)
			}
		}
		for id, n := range activeNotes {
			x, ok := n.Instrument.Value(coreTick, n, nil)
			if !ok {
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
	ID         uint64
	Instrument *WaveTable
	Tone       Tone
	Start      uint64
	End        uint64
	Attack     uint64
	Decay      uint64
	Sustain    float32
	Release    uint64
}
