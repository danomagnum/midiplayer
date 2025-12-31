package main

import "testing"

func TestSin(t *testing.T) {
	instrument := SinWave{}
	note := Note{Frequency: 440}
	var phase float64
	var results []float32
	for i := 0; i < 44100; i += 4410 {
		x := instrument.Value(phase, note, nil)
		results = append(results, x)
		phase += 4410
	}
	expected := []float32{0, 0.70710677, 1, 0.70710677, 0, -0.70710677, -1, -0.70710677, 0}
	for i, v := range expected {
		if float32Approx(results[i], v, 0.0001) == false {
			t.Errorf("At index %d, expected %f, got %f", i, v, results[i])
		}
	}
}

func float32Approx(a, b, tol float32) bool {
	if a > b {
		return a-b < tol
	}
	return b-a < tol
}
