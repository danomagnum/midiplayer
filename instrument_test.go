package main

import (
	"fmt"
	"testing"
)

func TestASDR(t *testing.T) {
	instr := NewInstrument("WaveForms/test.png", 10, 10, 0.1, 10, nil)
	note := Note{Start: 0, End: 30}

	for i := 0; i < 40; i++ {
		level := instr.adsr(uint64(i), note)
		fmt.Printf("Tick %d: Level %.2f\n", i, level)
	}
}
