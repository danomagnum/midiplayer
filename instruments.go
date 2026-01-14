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

func init() {
	dir, err := os.ReadDir("instruments")
	if err != nil {
		panic(err)
	}
	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		jsonFile, err := os.ReadFile("instruments/" + file.Name())
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
