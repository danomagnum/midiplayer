package main

import "math"

type Effect func(float32) float32
type EffectSetup interface {
	Effect(Note) Effect
}

type Reverb struct {
	rate     uint64
	channels int
	decay    float32
}

func (r *Reverb) Effect(note Note) Effect {

	history := make([]float32, r.rate)
	pos := 0
	channels := r.channels
	if r.channels == 0 {
		channels = 1
	}

	offsets := make([]int, channels)
	for i := 0; i < channels; i++ {
		offsets[i] = (i * len(history)) / channels
	}

	return func(in float32) float32 {
		p0 := pos
		new := float32(0)
		for i := 0; i < channels; i++ {
			new += history[(p0+offsets[i])%len(history)] * (r.decay / float32(channels))
		}
		out := in + new
		history[pos] = out
		pos = (pos + 1) % len(history)
		return out
	}
}

type LFO struct {
	Frequency float64
	Amplitude float64
}

func (lfo *LFO) Effect(note Note) Effect {
	return func(in float32) float32 {
		x := math.Sin(2 * math.Pi * float64(coreTick) * lfo.Frequency / baseRate)
		x = x*lfo.Amplitude + (1.0 - lfo.Amplitude)
		return in * float32(x)
	}
}
