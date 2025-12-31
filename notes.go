package main

// Chromatic note table (C0 to B8)
var (
	C0      = Tone{Frequency: 16.35, Name: "C0"}
	CSharp0 = Tone{Frequency: 17.32, Name: "C#0"}
	D0      = Tone{Frequency: 18.35, Name: "D0"}
	DSharp0 = Tone{Frequency: 19.45, Name: "D#0"}
	E0      = Tone{Frequency: 20.60, Name: "E0"}
	F0      = Tone{Frequency: 21.83, Name: "F0"}
	FSharp0 = Tone{Frequency: 23.12, Name: "F#0"}
	G0      = Tone{Frequency: 24.50, Name: "G0"}
	GSharp0 = Tone{Frequency: 25.96, Name: "G#0"}
	A0      = Tone{Frequency: 27.50, Name: "A0"}
	ASharp0 = Tone{Frequency: 29.14, Name: "A#0"}
	B0      = Tone{Frequency: 30.87, Name: "B0"}

	C1      = Tone{Frequency: 32.70, Name: "C1"}
	CSharp1 = Tone{Frequency: 34.65, Name: "C#1"}
	D1      = Tone{Frequency: 36.71, Name: "D1"}
	DSharp1 = Tone{Frequency: 38.89, Name: "D#1"}
	E1      = Tone{Frequency: 41.20, Name: "E1"}
	F1      = Tone{Frequency: 43.65, Name: "F1"}
	FSharp1 = Tone{Frequency: 46.25, Name: "F#1"}
	G1      = Tone{Frequency: 49.00, Name: "G1"}
	GSharp1 = Tone{Frequency: 51.91, Name: "G#1"}
	A1      = Tone{Frequency: 55.00, Name: "A1"}
	ASharp1 = Tone{Frequency: 58.27, Name: "A#1"}
	B1      = Tone{Frequency: 61.74, Name: "B1"}

	C2      = Tone{Frequency: 65.41, Name: "C2"}
	CSharp2 = Tone{Frequency: 69.30, Name: "C#2"}
	D2      = Tone{Frequency: 73.42, Name: "D2"}
	DSharp2 = Tone{Frequency: 77.78, Name: "D#2"}
	E2      = Tone{Frequency: 82.41, Name: "E2"}
	F2      = Tone{Frequency: 87.31, Name: "F2"}
	FSharp2 = Tone{Frequency: 92.50, Name: "F#2"}
	G2      = Tone{Frequency: 98.00, Name: "G2"}
	GSharp2 = Tone{Frequency: 103.83, Name: "G#2"}
	A2      = Tone{Frequency: 110.00, Name: "A2"}
	ASharp2 = Tone{Frequency: 116.54, Name: "A#2"}
	B2      = Tone{Frequency: 123.47, Name: "B2"}

	C3      = Tone{Frequency: 130.81, Name: "C3"}
	CSharp3 = Tone{Frequency: 138.59, Name: "C#3"}
	D3      = Tone{Frequency: 146.83, Name: "D3"}
	DSharp3 = Tone{Frequency: 155.56, Name: "D#3"}
	E3      = Tone{Frequency: 164.81, Name: "E3"}
	F3      = Tone{Frequency: 174.61, Name: "F3"}
	FSharp3 = Tone{Frequency: 185.00, Name: "F#3"}
	G3      = Tone{Frequency: 196.00, Name: "G3"}
	GSharp3 = Tone{Frequency: 207.65, Name: "G#3"}
	A3      = Tone{Frequency: 220.00, Name: "A3"}
	ASharp3 = Tone{Frequency: 233.08, Name: "A#3"}
	B3      = Tone{Frequency: 246.94, Name: "B3"}

	C4      = Tone{Frequency: 261.63, Name: "C4"}
	CSharp4 = Tone{Frequency: 277.18, Name: "C#4"}
	D4      = Tone{Frequency: 293.66, Name: "D4"}
	DSharp4 = Tone{Frequency: 311.13, Name: "D#4"}
	E4      = Tone{Frequency: 329.63, Name: "E4"}
	F4      = Tone{Frequency: 349.23, Name: "F4"}
	FSharp4 = Tone{Frequency: 369.99, Name: "F#4"}
	G4      = Tone{Frequency: 392.00, Name: "G4"}
	GSharp4 = Tone{Frequency: 415.30, Name: "G#4"}
	A4      = Tone{Frequency: 440.00, Name: "A4"}
	ASharp4 = Tone{Frequency: 466.16, Name: "A#4"}
	B4      = Tone{Frequency: 493.88, Name: "B4"}

	C5      = Tone{Frequency: 523.25, Name: "C5"}
	CSharp5 = Tone{Frequency: 554.37, Name: "C#5"}
	D5      = Tone{Frequency: 587.33, Name: "D5"}
	DSharp5 = Tone{Frequency: 622.25, Name: "D#5"}
	E5      = Tone{Frequency: 659.25, Name: "E5"}
	F5      = Tone{Frequency: 698.46, Name: "F5"}
	FSharp5 = Tone{Frequency: 739.99, Name: "F#5"}
	G5      = Tone{Frequency: 783.99, Name: "G5"}
	GSharp5 = Tone{Frequency: 830.61, Name: "G#5"}
	A5      = Tone{Frequency: 880.00, Name: "A5"}
	ASharp5 = Tone{Frequency: 932.33, Name: "A#5"}
	B5      = Tone{Frequency: 987.77, Name: "B5"}

	C6      = Tone{Frequency: 1046.50, Name: "C6"}
	CSharp6 = Tone{Frequency: 1108.73, Name: "C#6"}
	D6      = Tone{Frequency: 1174.66, Name: "D6"}
	DSharp6 = Tone{Frequency: 1244.51, Name: "D#6"}
	E6      = Tone{Frequency: 1318.51, Name: "E6"}
	F6      = Tone{Frequency: 1396.91, Name: "F6"}
	FSharp6 = Tone{Frequency: 1479.98, Name: "F#6"}
	G6      = Tone{Frequency: 1567.98, Name: "G6"}
	GSharp6 = Tone{Frequency: 1661.22, Name: "G#6"}
	A6      = Tone{Frequency: 1760.00, Name: "A6"}
	ASharp6 = Tone{Frequency: 1864.66, Name: "A#6"}
	B6      = Tone{Frequency: 1975.53, Name: "B6"}

	C7      = Tone{Frequency: 2093.00, Name: "C7"}
	CSharp7 = Tone{Frequency: 2217.46, Name: "C#7"}
	D7      = Tone{Frequency: 2349.32, Name: "D7"}
	DSharp7 = Tone{Frequency: 2489.02, Name: "D#7"}
	E7      = Tone{Frequency: 2637.02, Name: "E7"}
	F7      = Tone{Frequency: 2793.83, Name: "F7"}
	FSharp7 = Tone{Frequency: 2959.96, Name: "F#7"}
	G7      = Tone{Frequency: 3135.96, Name: "G7"}
	GSharp7 = Tone{Frequency: 3322.44, Name: "G#7"}
	A7      = Tone{Frequency: 3520.00, Name: "A7"}
	ASharp7 = Tone{Frequency: 3729.31, Name: "A#7"}
	B7      = Tone{Frequency: 3951.07, Name: "B7"}

	C8      = Tone{Frequency: 4186.01, Name: "C8"}
	CSharp8 = Tone{Frequency: 4434.92, Name: "C#8"}
	D8      = Tone{Frequency: 4698.63, Name: "D8"}
	DSharp8 = Tone{Frequency: 4978.03, Name: "D#8"}
	E8      = Tone{Frequency: 5274.04, Name: "E8"}
	F8      = Tone{Frequency: 5587.65, Name: "F8"}
	FSharp8 = Tone{Frequency: 5919.91, Name: "F#8"}
	G8      = Tone{Frequency: 6271.93, Name: "G8"}
	GSharp8 = Tone{Frequency: 6644.88, Name: "G#8"}
	A8      = Tone{Frequency: 7040.00, Name: "A8"}
	ASharp8 = Tone{Frequency: 7458.62, Name: "A#8"}
	B8      = Tone{Frequency: 7902.13, Name: "B8"}
)
