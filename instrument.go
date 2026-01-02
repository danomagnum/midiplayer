package main

type Instrument struct {
	Wave    WaveTable
	Attack  uint64
	Decay   uint64
	Sustain float32
	Release uint64
}

func NewInstrument(filename string, attack, decay uint64, sustain float32, release uint64) Instrument {
	wave := load(filename)
	return Instrument{
		Wave:    wave,
		Attack:  attack,
		Decay:   decay,
		Sustain: sustain,
		Release: release,
	}
}
