package main

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
