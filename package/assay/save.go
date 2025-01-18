package assay

import (
	"fmt"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

// ExportRawFFTToCSV exports FFT results to CSV. With float values: freq in Hz,
// magnitude, phase in rad.
func ExportRawFFTToCSV(fftResult []complex128, sampleRate float64, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for i, v := range fftResult {
		magnitude, phase := cmplx.Polar(v)
		frequency := float64(i) * sampleRate / float64(len(fftResult))
		_, err := fmt.Fprintf(file, "%.2f,%.5f,%.5f\n", frequency, magnitude, phase)
		if err != nil {
			return err
		}
	}

	return nil
}

// ExportFFTResultsToCSV exports extracted FFT results to CSV. With float values:
// freq in Hz, magnitude, phase in rad.
func ExportFFTResultsToCSV(frequencies, magnitudes, phases []float64, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for i, freq := range frequencies {
		_, err := fmt.Fprintf(file, "%.2f,%.5f,%.5f\n", freq, magnitudes[i], phases[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// PlotMagnitudeAndPhase generates two aligned plots: magnitude/frequency
// and phase/frequency, displayed one below the other.
func PlotMagnitudeAndPhase(frequencies, magnitudes, phases []float64, filename string) error {

	const rows, cols = 2, 1
	plots := make([][]*plot.Plot, rows)

	// Magnitude
	plots[0] = make([]*plot.Plot, cols)
	p, err := getMagnitudePlot(frequencies, magnitudes)
	if err != nil {
		return err
	}
	plots[0][0] = p

	// Phase
	plots[1] = make([]*plot.Plot, cols)
	p, err = getPhasePlot(frequencies, phases)
	if err != nil {
		return err
	}
	plots[1][0] = p

	// Draw
	img := vgimg.New(vg.Points(768), vg.Points(384))
	dc := draw.New(img)

	t := draw.Tiles{
		Rows: rows,
		Cols: cols,
	}

	canvases := plot.Align(plots, t, dc)
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			if plots[j][i] != nil {
				plots[j][i].Draw(canvases[j][i])
			}
		}
	}

	w, err := os.Create(filename)
	if err != nil {
		return err
	}

	png := vgimg.PngCanvas{Canvas: img}
	if _, err := png.WriteTo(w); err != nil {
		return err
	}

	return nil
}

func getMagnitudePlot(frequencies, magnitudes []float64) (*plot.Plot, error) {

	p := plot.New()
	p.Title.Text = "Magnitude Spectrum"
	p.X.Label.Text = "Frequency (Hz)"
	p.Y.Label.Text = "Magnitude"
	p.Y.Min = 0
	p.Y.Max = max(magnitudes) * 1.1

	magnitudePts := make(plotter.XYs, len(frequencies))
	for i := range frequencies {
		magnitudePts[i].X = frequencies[i]
		magnitudePts[i].Y = magnitudes[i]
	}

	magnitudeLine, err := plotter.NewLine(magnitudePts)
	if err != nil {
		return nil, err
	}
	magnitudeLine.Color = color.RGBA{B: 255, A: 255}
	p.Add(magnitudeLine)

	return p, nil
}

// TODO: ennek nincs értelme
func getPhasePlot(frequencies, phases []float64) (*plot.Plot, error) {

	p := plot.New()
	p.Title.Text = "Phase Spectrum"
	p.X.Label.Text = "Frequency (Hz)"
	p.Y.Label.Text = "Phase (radians)"
	p.Y.Min = -math.Pi
	p.Y.Max = math.Pi

	phasePts := make(plotter.XYs, len(frequencies))
	for i := range frequencies {
		phasePts[i].X = frequencies[i]
		phasePts[i].Y = phases[i]
	}

	phaseLine, err := plotter.NewLine(phasePts)
	if err != nil {
		return nil, err
	}
	phaseLine.Color = color.RGBA{R: 255, A: 255}
	p.Add(phaseLine)

	return p, nil
}

// PlotBins draws a diagram based on the values in the provided bins,
// with specified minimum and maximum values, and saves it to a file.
func PlotBins(bins []int, minValue, maxValue, numBins int, filename string) error {

	p := plot.New()
	p.Title.Text = "Distribution of Values in Bins"
	p.X.Label.Text = "Bins"
	p.Y.Label.Text = "Count"

	w := vg.Points(10)
	bars := make(plotter.Values, numBins)

	for i := 0; i < numBins; i++ {
		bars[i] = float64(bins[i])
	}

	barChart, err := plotter.NewBarChart(bars, w)
	if err != nil {
		return err
	}

	barChart.LineStyle.Width = vg.Length(0)
	barChart.Color = color.RGBA{R: 196, G: 128, B: 0, A: 255}

	binWidth := float64(maxValue-minValue) / float64(numBins)
	p.NominalX(makeBinLabels(numBins, binWidth, minValue, numBins/10)...)

	p.Add(barChart)

	img := vgimg.New(12*vg.Inch, 8*vg.Inch)
	dc := draw.New(img)
	p.Draw(dc)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := png.Encode(f, img.Image()); err != nil {
		return err
	}

	return nil
}

// labels for the bins in PlotBins at exact steps
func makeBinLabels(numBins int, binWidth float64, minValue int, step int) []string {
	labels := make([]string, numBins)

	for i := 0; i < numBins; i++ {
		if i%step == 0 {
			lowerBound := int(binWidth*float64(i)) + minValue
			upperBound := int(binWidth*float64(i+1)) + minValue - 1
			if i == numBins-1 {
				upperBound = minValue + int(binWidth*float64(numBins)) - 1
			}
			labels[i] = fmt.Sprintf("%d-%d", lowerBound, upperBound)
		} else {
			labels[i] = ""
		}
	}

	return labels
}
