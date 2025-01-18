package record

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"

	"github.com/youpy/go-wav"
)

// WriteToWav creates the given file and writes the buffer in WAV format.
func WriteToWav(sampleRate uint32, buffer []float64, fileName string) error {

	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("couldn't create file: '%s': %w", fileName, err)
	}
	defer f.Close()

	samplesConv := make([]wav.Sample, len(buffer))
	for idx, value := range buffer {
		samplesConv[idx].Values[0] = int(value * math.MaxInt32)
	}

	w := wav.NewWriter(f, uint32(len(samplesConv)), 1, sampleRate, 32)
	err = w.WriteSamples(samplesConv)
	if err != nil {
		return fmt.Errorf("couldn't write samples to file '%s': %w", fileName, err)
	}

	return nil
}

// WriteToRaw creates the given file and dumps the buffer in 32 bit integer format.
func WriteToRaw(buffer []float64, fileName string) error {

	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("couldn't create file: '%s': %w", fileName, err)
	}
	defer f.Close()

	var buf [4]byte
	for sampleIdx, sample := range buffer {

		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))

		_, err := f.Write(buf[:])
		if err != nil {
			return fmt.Errorf("couldn't write sample at index %d to file '%s': %w", sampleIdx, fileName, err)
		}
	}

	return nil
}
