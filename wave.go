package main

import (
	"image/png"
	"log"
	"log/slog"
	"os"
)

type Tone struct {
	Name      string
	Frequency float64
}

type Effect func(in float32) float32

type WaveTable struct {
	WaveForm []float32
}

func (s WaveTable) Value(tick uint64, note Note, effect Effect) (float32, bool) {
	// Sawtooth wave ranges from -1 to 1 over one period
	samples := len(s.WaveForm)
	period := baseRate / note.Tone.Frequency
	pos := tick % uint64(period)
	percent := float64(pos) / period
	index := int(percent * float64(samples))
	if index >= samples {
		index = samples - 1
	}
	x := s.WaveForm[index]

	adsr := note.Instrument.adsr(tick, note)
	x *= adsr
	lfo := note.Instrument.lfo(tick)
	x *= lfo

	cont := true
	if tick >= note.End+note.Instrument.Release {
		cont = false
	}
	slog.Debug("Note Playing", "tick", tick, "index", index, "adsr", adsr, "x", x)
	if effect == nil {
		return x, cont
	}
	return effect(x), cont
}

func load(path string) WaveTable {
	// Open the PNG file
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	// Decode the PNG image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatalf("failed to decode PNG: %v", err)
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	values := make([]float32, width)
	// Iterate across columns, bottom to top, and find first non-white pixel
	for x := 0; x < width; x++ {
		found := false
		for y := height - 1; y >= 0; y-- { // bottom to top
			clr := img.At(bounds.Min.X+x, bounds.Min.Y+y)
			r, g, b, a := clr.RGBA()
			// RGBA returns values in 16-bit (0-65535), white is all max and alpha max
			if !(r == 0xFFFF && g == 0xFFFF && b == 0xFFFF && a == 0xFFFF) {
				// Found first non-white pixel in this column
				// You can process/store (x, y) or color as needed here
				found = true
				values[x] = float32(y)
				break
			}
		}
		if !found {
			// Entire column is white
		}
	}

	// normalize values to -1.0 to 1.0
	minY := float32(0)
	maxY := float32(height - 1)
	for i := range values {
		if values[i] < minY {
			minY = values[i]
		}
		if values[i] > maxY {
			maxY = values[i]
		}
	}

	for i := range values {
		values[i] = ((values[i]-minY)/(maxY-minY))*2.0 - 1.0
	}

	return WaveTable{WaveForm: values}

}
