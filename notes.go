package main

// Chromatic note table (C0 to B8)
var (
	C0      = Note{Frequency: 16.35, Name: "C0"}
	CSharp0 = Note{Frequency: 17.32, Name: "C#0"}
	D0      = Note{Frequency: 18.35, Name: "D0"}
	DSharp0 = Note{Frequency: 19.45, Name: "D#0"}
	E0      = Note{Frequency: 20.60, Name: "E0"}
	F0      = Note{Frequency: 21.83, Name: "F0"}
	FSharp0 = Note{Frequency: 23.12, Name: "F#0"}
	G0      = Note{Frequency: 24.50, Name: "G0"}
	GSharp0 = Note{Frequency: 25.96, Name: "G#0"}
	A0      = Note{Frequency: 27.50, Name: "A0"}
	ASharp0 = Note{Frequency: 29.14, Name: "A#0"}
	B0      = Note{Frequency: 30.87, Name: "B0"}

	C1      = Note{Frequency: 32.70, Name: "C1"}
	CSharp1 = Note{Frequency: 34.65, Name: "C#1"}
	D1      = Note{Frequency: 36.71, Name: "D1"}
	DSharp1 = Note{Frequency: 38.89, Name: "D#1"}
	E1      = Note{Frequency: 41.20, Name: "E1"}
	F1      = Note{Frequency: 43.65, Name: "F1"}
	FSharp1 = Note{Frequency: 46.25, Name: "F#1"}
	G1      = Note{Frequency: 49.00, Name: "G1"}
	GSharp1 = Note{Frequency: 51.91, Name: "G#1"}
	A1      = Note{Frequency: 55.00, Name: "A1"}
	ASharp1 = Note{Frequency: 58.27, Name: "A#1"}
	B1      = Note{Frequency: 61.74, Name: "B1"}

	C2      = Note{Frequency: 65.41, Name: "C2"}
	CSharp2 = Note{Frequency: 69.30, Name: "C#2"}
	D2      = Note{Frequency: 73.42, Name: "D2"}
	DSharp2 = Note{Frequency: 77.78, Name: "D#2"}
	E2      = Note{Frequency: 82.41, Name: "E2"}
	F2      = Note{Frequency: 87.31, Name: "F2"}
	FSharp2 = Note{Frequency: 92.50, Name: "F#2"}
	G2      = Note{Frequency: 98.00, Name: "G2"}
	GSharp2 = Note{Frequency: 103.83, Name: "G#2"}
	A2      = Note{Frequency: 110.00, Name: "A2"}
	ASharp2 = Note{Frequency: 116.54, Name: "A#2"}
	B2      = Note{Frequency: 123.47, Name: "B2"}

	C3      = Note{Frequency: 130.81, Name: "C3"}
	CSharp3 = Note{Frequency: 138.59, Name: "C#3"}
	D3      = Note{Frequency: 146.83, Name: "D3"}
	DSharp3 = Note{Frequency: 155.56, Name: "D#3"}
	E3      = Note{Frequency: 164.81, Name: "E3"}
	F3      = Note{Frequency: 174.61, Name: "F3"}
	FSharp3 = Note{Frequency: 185.00, Name: "F#3"}
	G3      = Note{Frequency: 196.00, Name: "G3"}
	GSharp3 = Note{Frequency: 207.65, Name: "G#3"}
	A3      = Note{Frequency: 220.00, Name: "A3"}
	ASharp3 = Note{Frequency: 233.08, Name: "A#3"}
	B3      = Note{Frequency: 246.94, Name: "B3"}

	C4      = Note{Frequency: 261.63, Name: "C4"}
	CSharp4 = Note{Frequency: 277.18, Name: "C#4"}
	D4      = Note{Frequency: 293.66, Name: "D4"}
	DSharp4 = Note{Frequency: 311.13, Name: "D#4"}
	E4      = Note{Frequency: 329.63, Name: "E4"}
	F4      = Note{Frequency: 349.23, Name: "F4"}
	FSharp4 = Note{Frequency: 369.99, Name: "F#4"}
	G4      = Note{Frequency: 392.00, Name: "G4"}
	GSharp4 = Note{Frequency: 415.30, Name: "G#4"}
	A4      = Note{Frequency: 440.00, Name: "A4"}
	ASharp4 = Note{Frequency: 466.16, Name: "A#4"}
	B4      = Note{Frequency: 493.88, Name: "B4"}

	C5      = Note{Frequency: 523.25, Name: "C5"}
	CSharp5 = Note{Frequency: 554.37, Name: "C#5"}
	D5      = Note{Frequency: 587.33, Name: "D5"}
	DSharp5 = Note{Frequency: 622.25, Name: "D#5"}
	E5      = Note{Frequency: 659.25, Name: "E5"}
	F5      = Note{Frequency: 698.46, Name: "F5"}
	FSharp5 = Note{Frequency: 739.99, Name: "F#5"}
	G5      = Note{Frequency: 783.99, Name: "G5"}
	GSharp5 = Note{Frequency: 830.61, Name: "G#5"}
	A5      = Note{Frequency: 880.00, Name: "A5"}
	ASharp5 = Note{Frequency: 932.33, Name: "A#5"}
	B5      = Note{Frequency: 987.77, Name: "B5"}

	C6      = Note{Frequency: 1046.50, Name: "C6"}
	CSharp6 = Note{Frequency: 1108.73, Name: "C#6"}
	D6      = Note{Frequency: 1174.66, Name: "D6"}
	DSharp6 = Note{Frequency: 1244.51, Name: "D#6"}
	E6      = Note{Frequency: 1318.51, Name: "E6"}
	F6      = Note{Frequency: 1396.91, Name: "F6"}
	FSharp6 = Note{Frequency: 1479.98, Name: "F#6"}
	G6      = Note{Frequency: 1567.98, Name: "G6"}
	GSharp6 = Note{Frequency: 1661.22, Name: "G#6"}
	A6      = Note{Frequency: 1760.00, Name: "A6"}
	ASharp6 = Note{Frequency: 1864.66, Name: "A#6"}
	B6      = Note{Frequency: 1975.53, Name: "B6"}

	C7      = Note{Frequency: 2093.00, Name: "C7"}
	CSharp7 = Note{Frequency: 2217.46, Name: "C#7"}
	D7      = Note{Frequency: 2349.32, Name: "D7"}
	DSharp7 = Note{Frequency: 2489.02, Name: "D#7"}
	E7      = Note{Frequency: 2637.02, Name: "E7"}
	F7      = Note{Frequency: 2793.83, Name: "F7"}
	FSharp7 = Note{Frequency: 2959.96, Name: "F#7"}
	G7      = Note{Frequency: 3135.96, Name: "G7"}
	GSharp7 = Note{Frequency: 3322.44, Name: "G#7"}
	A7      = Note{Frequency: 3520.00, Name: "A7"}
	ASharp7 = Note{Frequency: 3729.31, Name: "A#7"}
	B7      = Note{Frequency: 3951.07, Name: "B7"}

	C8      = Note{Frequency: 4186.01, Name: "C8"}
	CSharp8 = Note{Frequency: 4434.92, Name: "C#8"}
	D8      = Note{Frequency: 4698.63, Name: "D8"}
	DSharp8 = Note{Frequency: 4978.03, Name: "D#8"}
	E8      = Note{Frequency: 5274.04, Name: "E8"}
	F8      = Note{Frequency: 5587.65, Name: "F8"}
	FSharp8 = Note{Frequency: 5919.91, Name: "F#8"}
	G8      = Note{Frequency: 6271.93, Name: "G8"}
	GSharp8 = Note{Frequency: 6644.88, Name: "G#8"}
	A8      = Note{Frequency: 7040.00, Name: "A8"}
	ASharp8 = Note{Frequency: 7458.62, Name: "A#8"}
	B8      = Note{Frequency: 7902.13, Name: "B8"}
)
