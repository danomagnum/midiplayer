package main

import (
	"encoding/json"
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
	LFO          []LFO
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
		var instrDef InstrumentDefinition
		err = json.Unmarshal(jsonFile, &instrDef)
		if err != nil {
			panic(err)
		}
		instr := NewInstrument(instrDef.WaveformPath, instrDef.Attack, instrDef.Decay, instrDef.Sustain, instrDef.Release, instrDef.LFO)
		DefaultInstruments[instrDef.ID] = instr

	}
}
