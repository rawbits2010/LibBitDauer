package player

import (
	"fmt"
	"io"
	"math"
	"time"

	"github.com/ebitengine/oto/v3"
)

const SoundOutputSampleRate = 44100
const SamplesPerMS = float64(SoundOutputSampleRate) * 1 / 1000.0

type SoundOutput struct {
	IsReady bool

	otoCtx           *oto.Context
	otoPlayer        *oto.Player
	playerDataReader PlayerDataReader
}

// NewSoundOutput creates and initialize a sound player that can output
// a 1 channel 44.1kHz sample composed of float64 values.
// There can only be one SoundOutput at any times!
func NewSoundOutput() (*SoundOutput, error) {
	outTmp := &SoundOutput{
		IsReady: false,
	}

	cop := &oto.NewContextOptions{
		ChannelCount: 1,
		SampleRate:   SoundOutputSampleRate,
		Format:       oto.FormatFloat32LE,
	}

	ctx, readyCh, err := oto.NewContext(cop)
	if err != nil {
		return nil, fmt.Errorf("couldn't initialize oto: %w", err)
	}

	outTmp.otoCtx = ctx
	go outTmp.createPlayerAfterReady(readyCh)

	return outTmp, nil
}

func (so *SoundOutput) createPlayerAfterReady(readyCh chan struct{}) {

	// need to wait for the hardware to initialize before creating a player
	<-readyCh

	so.playerDataReader = *NewPlayerDataReader()
	so.otoPlayer = so.otoCtx.NewPlayer(&so.playerDataReader)

	so.IsReady = true
}

// WaitUntilReady waits until the player is usable.
// Convenience for commandline usage.
func (so *SoundOutput) WaitUntilReady() {
	for {
		time.Sleep(time.Millisecond)
		if so.IsReady {
			break
		}
	}
}

// SetSamples will set the new samples to play.
// Currently it resets the sound playback to do so. It will resume playing if
// it was before.
func (so *SoundOutput) SetSamples(samples []float64) {

	isPlaying := so.IsPlaying()

	if isPlaying {
		so.otoPlayer.Pause()
	}

	so.playerDataReader.SetSamples(samples)
	if so.IsReady {
		so.otoPlayer.Seek(0, io.SeekStart)
	}

	if isPlaying {
		so.otoPlayer.Play()
	}
}

// Play starts the playback.
func (so *SoundOutput) Play() error {

	if !so.IsReady {
		return fmt.Errorf("player is not ready yet")
	}

	so.otoPlayer.Play()

	return nil
}

// Pause pauses the playback. Can be continued with Play.
func (so *SoundOutput) Pause() error {

	if !so.IsReady {
		return fmt.Errorf("player is not ready yet")
	}

	so.otoPlayer.Pause()

	return nil
}

// Stop stops the playback. Only needed if you want to stop playback early.
func (so *SoundOutput) Stop() error {

	if !so.IsReady {
		return fmt.Errorf("player is not ready yet")
	}

	so.otoPlayer.Pause()
	_, err := so.otoPlayer.Seek(0, io.SeekStart)
	if err != nil {
		panic(err) // should only happen when io.Seeker is not implemented
	}

	return nil
}

// IsPlaying returns if the player is playing audio.
func (so *SoundOutput) IsPlaying() bool {

	if !so.IsReady {
		return false
	}

	return so.otoPlayer.IsPlaying()
}

// WaitUntilStops waits until the player is usable.
// Convenience for commandline usage.
func (so *SoundOutput) WaitUntilStops() {
	for {
		time.Sleep(time.Millisecond)
		if !so.IsPlaying() {
			break
		}
	}
}

// SeekTo sets to playback position to corresponding sample position.
// It will return the timestamp for the actual position which can differ from
// the requested.
// This doesn't necessary mean an error.
func (so *SoundOutput) SeekTo(timestampMS float64) (float64, error) {

	if !so.IsReady {
		return 0, fmt.Errorf("player is not ready yet")
	}

	toSamplePos := int64(math.Ceil(SamplesPerMS * float64(timestampMS)))
	currSamplePos, err := so.otoPlayer.Seek(toSamplePos, io.SeekStart)
	if err != nil {
		panic(err) // should only happen when io.Seeker is not implemented
	}

	return float64(currSamplePos) / SamplesPerMS, nil
}

// Close closes the player. Call this with defer.
func (so *SoundOutput) Close() {
	so.IsReady = false
	if so.otoPlayer != nil {
		so.otoPlayer.Close()
	}
}
