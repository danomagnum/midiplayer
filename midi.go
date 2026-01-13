package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

type midiTrackData struct {
	events       map[int]map[int][]MIDIEvent
	InstrumentID int
	tempo        float64
	timeDivision uint16
	notes        []Note
}

func (m *midiTrackData) scaleFactor() float64 {

	tempo := 60.0 / m.tempo
	return (baseRate) * tempo / float64(m.timeDivision)
}

func LoadMIDIFile(path string) ([]Note, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	time := uint32(0)

	var header MIDIHeader
	var songData midiTrackData
	songData.events = make(map[int]map[int][]MIDIEvent)
	songData.tempo = 120.0 // Default tempo

	for {
		chunkType := make([]byte, 4)
		_, err := file.Read(chunkType)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		chunkLength := uint32(0)
		err = binary.Read(file, binary.BigEndian, &chunkLength)
		if err != nil {
			return nil, err
		}
		chunkData := make([]byte, chunkLength)
		_, err = file.Read(chunkData)
		if err != nil {
			return nil, err
		}
		chunkBuffer := bytes.NewBuffer(chunkData)

		//tracks := []MIDITrack{}

		switch string(chunkType) {
		case "MThd":
			if chunkLength != 6 {
				return nil, fmt.Errorf("invalid MIDI header length: %d", chunkLength)
			}
			header.FormatType = binary.BigEndian.Uint16(chunkData[0:2])
			header.NumberOfTracks = binary.BigEndian.Uint16(chunkData[2:4])
			header.TimeDivision = binary.BigEndian.Uint16(chunkData[4:6])
			songData.timeDivision = header.TimeDivision
			fmt.Printf("MIDI Header: FormatType=%d, NumberOfTracks=%d, TimeDivision=%d\n",
				header.FormatType, header.NumberOfTracks, header.TimeDivision)

			if header.TimeDivision&0x8000 == 0 {
				// ticks per quarter
				ticksPerQuarter := header.TimeDivision & 0x7FFF
				songData.timeDivision = header.TimeDivision
				fmt.Printf("Ticks per Quarter Note: %d\n", ticksPerQuarter)
			} else {
				negSMPTE := header.TimeDivision & 0x7F00
				ticksPerFrame := header.TimeDivision & 0x00FF
				fmt.Printf("SMPTE Format: %d, Ticks per Frame: %d\n", negSMPTE, ticksPerFrame)

			}
		case "MTrk":
			fmt.Println("Found MIDI Track Chunk")
			for {
				deltaTime, err := readVariableLengthQuantity(chunkBuffer)
				if err != nil {
					if err.Error() == "EOF" {
						break
					}
					return nil, err
				}
				time += deltaTime

				midiMessage, err := chunkBuffer.ReadByte()
				if err != nil {
					return nil, err
				}
				if midiMessage == 0xFF {
					// Meta Event
					metaType, err := chunkBuffer.ReadByte()
					if err != nil {
						return nil, err
					}
					length, err := readVariableLengthQuantity(chunkBuffer)
					if err != nil {
						return nil, err
					}
					data := make([]byte, length)
					_, err = chunkBuffer.Read(data)
					if err != nil {
						return nil, err
					}
					fmt.Printf("Meta Event - Type: 0x%X, Length: %d, Data: %v\n", metaType, length, data)
					continue
				}
				msgType := midiMessage & 0xF0
				channel := midiMessage & 0x0F
				switch msgType {
				case 0x90:
					// Note On
					noteData := make([]byte, 2)
					_, err = chunkBuffer.Read(noteData)
					if err != nil {
						return nil, err
					}
					noteNumber := noteData[0]
					velocity := noteData[1]

					cd, ok := songData.events[int(channel)]
					if !ok {
						cd = make(map[int][]MIDIEvent)
						songData.events[int(channel)] = cd
					}
					nd, ok := cd[int(noteNumber)]
					if !ok {
						nd = make([]MIDIEvent, 0)
						cd[int(noteNumber)] = nd
					}
					if velocity == 0 {
						// Treat as Note Off
						fmt.Printf("Note Off - Channel: %d, Note: %d (velocity 0)\n", channel, noteNumber)
						songData.events[int(channel)][int(noteNumber)] = append(songData.events[int(channel)][int(noteNumber)], MIDIEvent{
							DeltaTime: uint32(float64(time) * songData.scaleFactor()),
							OnMessage: false,
							Velocity:  int(velocity),
						})
					} else {
						fmt.Printf("Note On - Channel: %d, Note: %d, Velocity: %d\n", channel, noteNumber, velocity)
						songData.events[int(channel)][int(noteNumber)] = append(songData.events[int(channel)][int(noteNumber)], MIDIEvent{
							DeltaTime: uint32(float64(time) * songData.scaleFactor()),
							OnMessage: true,
							Velocity:  int(velocity),
						})
					}
				case 0x80:
					// Note Off
					noteData := make([]byte, 2)
					_, err = chunkBuffer.Read(noteData)
					if err != nil {
						return nil, err
					}
					noteNumber := noteData[0]
					velocity := noteData[1]
					fmt.Printf("Note Off - Channel: %d, Note: %d, Velocity: %d\n", channel, noteNumber, velocity)

					songData.events[int(channel)][int(noteNumber)] = append(songData.events[int(channel)][int(noteNumber)], MIDIEvent{
						DeltaTime: uint32(float64(time) * songData.scaleFactor()),
						OnMessage: false,
						Velocity:  int(velocity),
					})

				case 0xC0:
					// Program Change
					progData := make([]byte, 1)
					_, err = chunkBuffer.Read(progData)
					if err != nil {
						return nil, err
					}
					programNumber := progData[0]
					fmt.Printf("Program Change - Channel: %d, Program: %d\n", channel, programNumber)
				default:
					fmt.Printf("Unhandled MIDI Message Type: 0x%X on Channel %d\n", msgType, channel)
				}
				fmt.Printf("Delta Time: %d, MIDI Message: 0x%X\n", deltaTime, midiMessage)
			}

		default:
			fmt.Println("Unknown Chunk Type:", string(chunkType))
			return nil, fmt.Errorf("unknown chunk type: %s", string(chunkType))
		}
	}

	notes := []Note{}
	for noteID, noteEvents := range songData.events[0] {
		var currentNote Note
		currentNote.Tone = MidiNoteToTone(noteID)
		for _, event := range noteEvents {
			if event.OnMessage {
				currentNote.Start = uint64(event.DeltaTime)
			} else {
				currentNote.End = uint64(event.DeltaTime)
				notes = append(notes, currentNote)
			}
		}

	}
	return notes, nil
}

type MIDIHeader struct {
	FormatType     uint16
	NumberOfTracks uint16
	TimeDivision   uint16
}

func readVariableLengthQuantity(data *bytes.Buffer) (uint32, error) {
	var value uint32
	for {
		b, err := data.ReadByte()
		if err != nil {
			return 0, err
		}
		value = (value << 7) | uint32(b&0x7F)
		if b&0x80 == 0 {
			break
		}
	}
	return value, nil

}

type MIDIEvent struct {
	DeltaTime uint32
	OnMessage bool
	Velocity  int
}

func MidiNoteToTone(noteNumber int) Tone {
	noteNames := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	octave := (noteNumber / 12) - 1
	noteIndex := noteNumber % 12
	noteName := fmt.Sprintf("%s%d", noteNames[noteIndex], octave)
	frequency := 440.0 * math.Pow(2, float64(noteNumber-69)/12.0)
	return Tone{
		Name:      noteName,
		Frequency: frequency,
	}
}
