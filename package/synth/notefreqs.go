package synth

import "math"

// NOTE: only 12 note system is supported - microtonal stuff maybe later
var freqTable [108]float64

// GenerateFreqTableBasedOn generates a 12 note LUT based on an arbitrary
// frequency in Hz used as the A note of the 4th octave. Can't use the LUT
// before this.
// This value changed a lot throughout history. Also, some instruments are
// tuned differently - currently only 1 LUT is supported. Use GetFreqOf to
// get the frequency of a note.
func GenerateFreqTableBasedOn(freqOfA4 float64) {
	for octave := 0; octave <= 8; octave++ {
		for note := 0; note < 12; note++ {
			noteIdx := octave*12 + note
			freqTable[noteIdx] = freqOfA4 * math.Pow(2, (float64(noteIdx-8)-49)/12)
		}
	}
}

const (
	NoteC    uint8 = 0
	NoteCisz uint8 = 1
	NoteDesz uint8 = 1
	NoteD    uint8 = 2
	NoteDisz uint8 = 3
	NoteEsz  uint8 = 3
	NoteE    uint8 = 4
	NoteF    uint8 = 5
	NoteFisz uint8 = 6
	NoteGesz uint8 = 6
	NoteG    uint8 = 7
	NoteGisz uint8 = 8
	NoteAsz  uint8 = 8
	NoteA    uint8 = 9
	NoteAisz uint8 = 10
	NoteB    uint8 = 10
	NoteH    uint8 = 11
)

// GetFreqOf returns the frequency of a note from the pre set LUT. Use
// GenerateFreqTableBasedOn to generate it.
func GetFreqOf(octave uint8, note uint8) float64 {
	return freqTable[octave*12+note]
}
