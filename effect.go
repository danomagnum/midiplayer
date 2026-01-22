package main

import "math"

type Effect func(float32) float32
type EffectSetup interface {
	Effect(Note) Effect
	ApplicationOrder() int
}

type Reverb struct {
	Rate     uint64
	Channels int
	Decay    float32
	Order    int
}

func (r *Reverb) ApplicationOrder() int {
	return r.Order
}

func (r *Reverb) Effect(note Note) Effect {

	history := make([]float32, r.Rate)
	pos := 0
	channels := r.Channels
	if r.Channels == 0 {
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
			new += history[(p0+offsets[i])%len(history)] * (r.Decay / float32(channels))
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
	Order     int
}

func (lfo *LFO) ApplicationOrder() int {
	return lfo.Order
}

func (lfo *LFO) Effect(note Note) Effect {
	return func(in float32) float32 {
		x := math.Sin(2 * math.Pi * float64(coreTick) * lfo.Frequency / baseRate)
		x = x*lfo.Amplitude + (1.0 - lfo.Amplitude)
		return in * float32(x)
	}
}

type Distortion struct {
	Gain  float32
	Order int
}

func (d *Distortion) Effect(note Note) Effect {
	return func(in float32) float32 {
		out := in * d.Gain
		if out > 1.0 {
			out = 1.0
		} else if out < -1.0 {
			out = -1.0
		}
		return out
	}
}

func (d *Distortion) ApplicationOrder() int {
	return d.Order
}
