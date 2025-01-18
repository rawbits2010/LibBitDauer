# LibBitDauer v0.1

A simple software synthesizer package for my audio projects. Also includes some basic sound analysis tools for no reason.

#### Limitations:
- It has some limited support for continuous generation, but currently it is only used to generate into a fixed size buffer.
- The envelope is also supposed to fully support the MIDI key-trigger mechanics, but it's never been used.

## Features
- Oscillator
  - Wave functions:
    - Sine, square, triangle, sawtooth, reverse sawtooth
	- Noise:
	  - White (using built-in Go rand.Rand)
	  - Red, pink, blue, violet (by applying filters to the white noise)
- Modulators:
  - Envelope (ADSR with various shapes):
    - LERP
    - Ease-in, ease-out, ease-in-out
	- Exponential, logarithmic, inverse exponential, inverse logarithmic
	- S-Curve (sigmoid)
  - LFO (can use all wave functions above)
- Filters:
  - Single, chain, chain of chain
  - Low-pass, high-pass, band-pass, notch, peaking
- Musical scale LUT generator based on a base note frequency
- Output options:
  - Export to WAV or raw (unsigned 32 bit integer) format
  - Play as 1 channel 44.1kHz using the [oto package](https://github.com/ebitengine/oto)

## Examples
Some examples are provided in the cmd folder, both for the synth and the analysis tools.

### Genex
The synth example creates an oscillator that outputs an A4 note with a customized envelope into a 2-second-long buffer. The generated sound will be played with the default sound device and saved as a WAV file.

### Idex
Loads a WAV file and extracts the dominant frequencies to identify the waveform. It also saves the extracted data into a CSV, and generates various diagrams into PNGs. An example untuned identification is performed using the harmonic ratios and another with other metrics (bandwidth, spectral centroid).

## TODOs
A rough list of planned features:

- Mixer to have multiple oscillators
- PWM generator
- Filter support to the oscillator
- Free form modulator (multiple envelope sections with selectable easing functions)
- Other filters like expander, compressor, limiter
- Handle overdrive

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.