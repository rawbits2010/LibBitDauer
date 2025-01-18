package player

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type PlayerDataReader struct {
	samples []float64
	pos     int
}

// NewPlayerDataReader creates a reader object that is used by the player to play the samples.
// You don't need to do anything with this construct, it is for the player.
func NewPlayerDataReader() *PlayerDataReader {
	return &PlayerDataReader{samples: nil, pos: 0}
}

// SetSamples switches out the samples to play. It is a thread-safe operation.
// It won't change the current player's position, there is the Seek function for that.
func (p *PlayerDataReader) SetSamples(samples []float64) {
	p.samples = samples
}

func (p *PlayerDataReader) Read(buff []byte) (int, error) {

	// 4 bytes per sample, 1 channel
	var tmp [4]byte
	var n = 0
	for ; n < int(len(buff)/4); n++ {
		if p.pos+n >= len(p.samples) {
			break
		}

		binary.LittleEndian.PutUint32(tmp[:], math.Float32bits(float32(p.samples[p.pos+n])))
		copy(buff[4*n:], tmp[:])
	}

	p.pos += n

	if n == 0 {
		return 0, io.EOF
	}

	return n * 4, nil
}

func (p *PlayerDataReader) Seek(offset int64, whence int) (int64, error) {

	var newPos int
	switch whence {
	case io.SeekStart:
		newPos = int(offset)
	case io.SeekCurrent:
		newPos = p.pos + int(offset)
	case io.SeekEnd:
		newPos = p.pos + int(offset)
	default:
		return 0, fmt.Errorf("invalid relative seek position (whence) '%d'", whence)
	}

	// clamping is better than an error (and more meaningful with returning the actual new position)
	if newPos < 0 {
		newPos = 0
	}
	if newPos > len(p.samples) {
		newPos = len(p.samples)
	}

	p.pos = newPos

	return int64(newPos), nil
}
