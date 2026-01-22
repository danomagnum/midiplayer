package main

import (
	"encoding/json"
	"log/slog"
	"os"
	"strings"
)

type InstrumentDefinition struct {
	Name         string
	ID           int
	WaveformPath string
	Attack       uint64
	Decay        uint64
	Sustain      float32
	Release      uint64
	LFOs         []LFO
}

var DefaultInstruments = map[int]Instrument{}

// TODO: eventually we'll want to go:embed these files both the json definitions and the waveform pngs
func init() {
	loadInstrumentsFromDir("instruments")
}

func loadInstrumentsFromDir(dirname string) {
	dir, err := os.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	for _, file := range dir {
		if file.IsDir() {
			loadInstrumentsFromDir(dirname + "/" + file.Name())
			continue
		}
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		jsonFile, err := os.ReadFile(dirname + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		var instrDef []InstrumentDefinition
		err = json.Unmarshal(jsonFile, &instrDef)
		if err != nil {
			slog.Warn("Bad json file", "name", file.Name(), "dir", dirname, "error", err)
		}
		for _, def := range instrDef {
			instr := NewInstrument(def.WaveformPath, def.Attack, def.Decay, def.Sustain, def.Release, def.LFOs)
			DefaultInstruments[def.ID] = instr
		}
	}

	piano := DefaultInstruments[1]
	piano.Effects = append(piano.Effects, &Reverb{rate: 10000, decay: 1.0, channels: 5})
	DefaultInstruments[2] = piano

}
