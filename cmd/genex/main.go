package main

import (
	"math"

	"github.com/rawbits2010/LibBitDauer/package/player"
	"github.com/rawbits2010/LibBitDauer/package/record"
	"github.com/rawbits2010/LibBitDauer/package/synth"
	"github.com/rawbits2010/LibBitDauer/package/synth/buffer"
	"github.com/rawbits2010/LibBitDauer/package/synth/generator"
	"github.com/rawbits2010/LibBitDauer/package/synth/modulator/easing"
)

func main() {

	player, err := player.NewSoundOutput()
	if err != nil {
		panic(err)
	}
	defer player.Close()

	player.WaitUntilReady()

	duration := uint(2)
	sampleRate := uint(44100)

	synth.GenerateFreqTableBasedOn(440)

	osc := synth.NewOscillator(sampleRate)
	osc.SwitchGeneratorType(synth.OscTypeWave)
	osc.Wave.SetFunction(generator.SawtoothFunction)
	osc.Volume = 1
	osc.Frequency = synth.GetFreqOf(4, synth.NoteA)

	osc.Envelope.AttackCurve = easing.NewExponential(10)
	osc.Envelope.ReleaseCurve = easing.NewLogarithmic(math.E)

	osc.Envelope.SetAttackLength(duration / 2 * 1000)
	osc.Envelope.SetDecayLength(0)
	osc.Envelope.SetSustain(1)
	osc.Envelope.SetReleaseLength(duration / 2 * 1000)
	osc.UseEnvelope = true

	buff := buffer.Generate(duration*1000, osc)

	player.SetSamples(buff)
	player.Play()
	player.WaitUntilStops()

	record.WriteToWav(uint32(sampleRate), buff, "out.wav")
}
