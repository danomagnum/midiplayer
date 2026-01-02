package main

import "math"

type LFO struct {
	Frequency float64
	Amplitude float64
}

type Instrument struct {
	Wave    WaveTable
	Attack  uint64
	Decay   uint64
	Sustain float32
	Release uint64
	LFO     []LFO
}

func NewInstrument(filename string, attack, decay uint64, sustain float32, release uint64, lfo []LFO) Instrument {
	wave := load(filename)
	return Instrument{
		Wave:    wave,
		Attack:  attack,
		Decay:   decay,
		Sustain: sustain,
		Release: release,
		LFO:     lfo,
	}
}

func (inst *Instrument) adsr(tick uint64, note Note) float32 {
	pos := tick - note.Start
	switch {
	case pos < inst.Attack:
		// Attack phase
		attack_level := float32(pos) / float32(inst.Attack)
		return attack_level
	case pos < inst.Attack+inst.Decay:
		// Decay phase
		decay_pos := pos - inst.Attack
		decay_level := 1.0 - (1.0-float32(inst.Sustain))*float32(decay_pos)/float32(inst.Decay)
		return decay_level
	case tick < note.End:
		// Sustain phase
		return inst.Sustain
	case tick >= note.End:
		// Release phase
		release_pos := tick - note.End
		release_level := 1.0 - float32(release_pos)/float32(inst.Release)
		if release_level < 0 {
			release_level = 0
		}
		return release_level
	default:
		return 0.0
	}

}

func (inst *Instrument) lfo(tick uint64) float32 {
	if len(inst.LFO) == 0 {
		return 1.0
	}
	out := 1.0
	for _, v := range inst.LFO {

		x := math.Sin(2 * math.Pi * float64(tick) * v.Frequency / baseRate)
		// scale from amplitude to 1.0
		x = x*v.Amplitude + (1.0 - v.Amplitude)
		out *= x

	}
	return float32(out)
}
