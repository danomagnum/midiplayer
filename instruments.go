package main

var instrument1 = NewInstrument("WaveForms/test.png", 1000, 1000, 0.8, 2000, nil)
var fifth = NewInstrument("WaveForms/fifth.png", 1000, 1000, 0.8, 2000, nil)
var piano = NewInstrument("WaveForms/piano.png", 800, 800, 0.4, 200, nil)
var violin = NewInstrument("WaveForms/violin.png", 2000, 0, 1, 3000, nil)
var organ = NewInstrument("WaveForms/organ.png", 4000, 0, 1, 4000, []LFO{
	{Frequency: 5.0, Amplitude: 0.1},
	{Frequency: 7.0, Amplitude: 0.1},
})
var trumpet = NewInstrument("WaveForms/trumpet.png", 2000, 2000, 0.5, 3000, nil)

// MIDI instruments (48-55)
var stringEnsemble1 = NewInstrument("WaveForms/generated/string_ensemble_1.png", 2000, 1000, 0.7, 3000, nil) // Sawtooth, lush
var stringEnsemble2 = NewInstrument("WaveForms/generated/string_ensemble_2.png", 1500, 800, 0.6, 2500, nil)  // Sawtooth, brighter
var synthStrings1 = NewInstrument("WaveForms/generated/synthstrings_1.png", 1000, 1000, 0.8, 2000, nil)      // Synthetic saw
var synthStrings2 = NewInstrument("WaveForms/generated/synthstrings_2.png", 800, 800, 0.7, 1500, nil)        // Square/saw hybrid
var choirAahs = NewInstrument("WaveForms/generated/choir_aahs.png", 3000, 2000, 0.9, 4000, nil)              // Smooth sine/triangle
var voiceOohs = NewInstrument("WaveForms/generated/voice_oohs.png", 3500, 2000, 0.8, 3500, nil)              // Softer sine/triangle
var synthVoice = NewInstrument("WaveForms/generated/synth_voice.png", 1200, 1000, 0.7, 2000, nil)            // Synthetic triangle/sine
var orchestraHit = NewInstrument("WaveForms/generated/orchestra_hit.png", 100, 100, 0.5, 500, nil)           // Sharp, percussive

// MIDI instruments 2-8
var brightAcousticPiano = NewInstrument("WaveForms/generated/bright_acoustic_piano.png", 800, 800, 0.4, 200, nil) // Brighter piano
var electricGrandPiano = NewInstrument("WaveForms/generated/electric_grand_piano.png", 700, 700, 0.5, 200, nil)   // Electric grand
var honkyTonkPiano = NewInstrument("WaveForms/generated/honky_tonk_piano.png", 900, 900, 0.3, 150, nil)           // Honky tonk
var electricPiano1 = NewInstrument("WaveForms/generated/electric_piano_1.png", 1000, 1000, 0.5, 300, nil)         // Rhodes
var electricPiano2 = NewInstrument("WaveForms/generated/electric_piano_2.png", 1000, 1000, 0.5, 300, nil)         // Wurlitzer
var harpsichord = NewInstrument("WaveForms/generated/harpsichord.png", 200, 200, 0.7, 100, nil)                   // Harpsichord
var clavinet = NewInstrument("WaveForms/generated/clavinet.png", 100, 100, 0.6, 80, nil)                          // Clavinet

// MIDI instruments 9-16
var celesta = NewInstrument("WaveForms/generated/celesta.png", 200, 200, 0.7, 100, nil)            // Celesta, bell-like
var glockenspiel = NewInstrument("WaveForms/generated/glockenspiel.png", 100, 100, 0.6, 80, nil)   // Glockenspiel, bright bell
var musicBox = NewInstrument("WaveForms/generated/music_box.png", 150, 150, 0.5, 100, nil)         // Music Box, plucked bell
var vibraphone = NewInstrument("WaveForms/generated/vibraphone.png", 300, 300, 0.7, 200, nil)      // Vibraphone, mellow bell
var marimba = NewInstrument("WaveForms/generated/marimba.png", 200, 200, 0.6, 120, nil)            // Marimba, wood
var xylophone = NewInstrument("WaveForms/generated/xylophone.png", 100, 100, 0.5, 80, nil)         // Xylophone, bright wood
var tubularBells = NewInstrument("WaveForms/generated/tubular_bells.png", 400, 400, 0.8, 300, nil) // Tubular Bells, deep bell
var dulcimer = NewInstrument("WaveForms/generated/dulcimer.png", 250, 250, 0.6, 120, nil)          // Dulcimer, plucked string

// MIDI instruments 17-24
var drawbarOrgan = NewInstrument("WaveForms/generated/drawbar_organ.png", 1000, 0, 1, 2000, []LFO{{Frequency: 5.0, Amplitude: 0.08}, {Frequency: 7.0, Amplitude: 0.05}}) // Drawbar organ, LFO for vibrato
var percussiveOrgan = NewInstrument("WaveForms/generated/percussive_organ.png", 400, 200, 0.7, 800, []LFO{{Frequency: 5.0, Amplitude: 0.04}})                            // Percussive organ, light vibrato
var rockOrgan = NewInstrument("WaveForms/generated/rock_organ.png", 600, 200, 0.8, 1000, []LFO{{Frequency: 6.0, Amplitude: 0.07}})                                       // Rock organ, moderate vibrato
var churchOrgan = NewInstrument("WaveForms/generated/church_organ.png", 2000, 0, 1, 4000, []LFO{{Frequency: 5.0, Amplitude: 0.03}})                                      // Church organ, slow vibrato
var reedOrgan = NewInstrument("WaveForms/generated/reed_organ.png", 800, 200, 0.7, 1200, []LFO{{Frequency: 5.5, Amplitude: 0.06}})                                       // Reed organ, vibrato
var accordion = NewInstrument("WaveForms/generated/accordion.png", 700, 300, 0.8, 1000, []LFO{{Frequency: 6.5, Amplitude: 0.08}})                                        // Accordion, vibrato
var harmonica = NewInstrument("WaveForms/generated/harmonica.png", 300, 100, 0.7, 600, []LFO{{Frequency: 5.5, Amplitude: 0.09}})                                         // Harmonica, strong vibrato
var bandoneon = NewInstrument("WaveForms/generated/bandoneon.png", 700, 300, 0.8, 1000, []LFO{{Frequency: 6.0, Amplitude: 0.07}})                                        // Bandoneon, vibrato

// MIDI instruments 25-32 (Guitars)
var acousticGuitarNylon = NewInstrument("WaveForms/generated/acoustic_guitar_nylon.png", 60, 80, 0.85, 180, nil)                                     // Nylon, punchier
var acousticGuitarSteel = NewInstrument("WaveForms/generated/acoustic_guitar_steel.png", 40, 60, 0.85, 150, nil)                                     // Steel, punchier
var electricGuitarJazz = NewInstrument("WaveForms/generated/electric_guitar_jazz.png", 30, 60, 0.8, 120, nil)                                        // Jazz, punchier
var electricGuitarClean = NewInstrument("WaveForms/generated/electric_guitar_clean.png", 20, 40, 0.8, 100, []LFO{{Frequency: 5.5, Amplitude: 0.04}}) // Clean, punchier
var electricGuitarMuted = NewInstrument("WaveForms/generated/electric_guitar_muted.png", 10, 20, 0.4, 40, nil)                                       // Muted, very percussive
var overdrivenGuitar = NewInstrument("WaveForms/generated/overdriven_guitar.png", 10, 20, 0.7, 60, nil)                                              // Overdriven, punchier
var distortionGuitar = NewInstrument("WaveForms/generated/distortion_guitar.png", 10, 20, 0.6, 60, nil)                                              // Distortion, punchier
var guitarHarmonics = NewInstrument("WaveForms/generated/guitar_harmonics.png", 30, 40, 0.85, 100, nil)                                              // Harmonics, punchier

var DefaultInstruments = map[int]Instrument{
	1:  piano,
	2:  brightAcousticPiano,
	3:  electricGrandPiano,
	4:  honkyTonkPiano,
	5:  electricPiano1,
	6:  electricPiano2,
	7:  harpsichord,
	8:  clavinet,
	9:  celesta,
	10: glockenspiel,
	11: musicBox,
	12: vibraphone,
	13: marimba,
	14: xylophone,
	15: tubularBells,
	16: dulcimer,
	17: organ,
	//17: drawbarOrgan,
	18: percussiveOrgan,
	19: rockOrgan,
	20: churchOrgan,
	21: reedOrgan,
	22: accordion,
	23: harmonica,
	24: bandoneon,
	25: acousticGuitarNylon,
	26: acousticGuitarSteel,
	27: electricGuitarJazz,
	28: electricGuitarClean,
	29: electricGuitarMuted,
	30: overdrivenGuitar,
	31: distortionGuitar,
	32: guitarHarmonics,
	40: violin,
	48: stringEnsemble1,
	49: stringEnsemble2,
	50: synthStrings1,
	51: synthStrings2,
	52: choirAahs,
	53: voiceOohs,
	54: synthVoice,
	55: orchestraHit,
	56: trumpet,
}
