package main

import "math"

type SinWave struct{}

func (s SinWave) Value(tick uint64, note Tone, effect Effect) float32 {
	x := float32(math.Sin(2 * math.Pi * float64(tick) * note.Frequency / baseRate))
	if effect == nil {
		return x
	}
	return effect(x)
}

type SawTooth struct{}

func (s SawTooth) Value(tick uint64, note Tone, effect Effect) float32 {
	// Sawtooth wave ranges from -1 to 1 over one period
	period := baseRate / note.Frequency
	pos := tick % uint64(period)
	x := float32((2.0 * (float64(pos) / period)) - 1.0)
	if effect == nil {
		return x
	}
	return effect(x)
}

type SquareWave struct {
	DutyCycle float64 // 0.0 to 1.0
}

func (s SquareWave) Value(tick uint64, note Tone, effect Effect) float32 {
	// Sawtooth wave ranges from -1 to 1 over one period
	period := baseRate / note.Frequency
	pos := tick % uint64(period)
	x := float32(-1.0)
	if float64(pos) < s.DutyCycle*period {
		x = 1.0
	}
	if effect == nil {
		return x
	}
	return effect(x)
}
