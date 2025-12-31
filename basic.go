package main

import "math"

type SinWave struct{}

func (s SinWave) Value(phase float64, note Note, effect Effect) float32 {
	x := float32(math.Sin(2 * math.Pi * phase * note.Frequency / 44100))
	if effect == nil {
		return x
	}
	return effect(x)
}

type SawTooth struct{}

func (s SawTooth) Value(phase float64, note Note, effect Effect) float32 {
	// Sawtooth wave ranges from -1 to 1 over one period
	period := 44100 / note.Frequency
	pos := int(phase) % int(period)
	x := float32((2.0 * (float64(pos) / period)) - 1.0)
	if effect == nil {
		return x
	}
	return effect(x)
}

type SquareWave struct {
	DutyCycle float64 // 0.0 to 1.0
}

func (s SquareWave) Value(phase float64, note Note, effect Effect) float32 {
	// Sawtooth wave ranges from -1 to 1 over one period
	period := 44100 / note.Frequency
	pos := int(phase) % int(period)
	x := float32(-1.0)
	if float64(pos) < s.DutyCycle*period {
		x = 1.0
	}
	if effect == nil {
		return x
	}
	return effect(x)
}
