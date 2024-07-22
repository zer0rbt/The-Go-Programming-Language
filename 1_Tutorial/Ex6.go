// Exercise 1.6: Modify the Lissajous program to produce images in multiple colors by adding
//more values to palette and then displaying them by changing the third argument of SetColorIndex in some interesting way.

// Lissajous generates GIF animations of random Lissajous figures.

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var startColor = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}
var palette []color.Color

func semiRandColor(initColor color.RGBA, diff int) color.RGBA {
	rand.Seed(time.Now().UnixNano())

	randomInRange := func(diff int) int {
		return rand.Intn(2*diff+1) - diff
	}

	randR := randomInRange(diff)
	randG := randomInRange(diff)
	randB := randomInRange(diff)

	clamp := func(value int) uint8 {
		if value < 0 {
			return 0
		}
		if value > 255 {
			return 255
		}
		return uint8(value)
	}

	newR := clamp(int(initColor.R) + randR)
	newG := clamp(int(initColor.G) + randG)
	newB := clamp(int(initColor.B) + randB)

	return color.RGBA{R: newR, G: newG, B: newB, A: initColor.A}
}

func generatePalette(length int, diff int) []color.Color {
	newPalette := make([]color.Color, length)
	newPalette[0] = startColor
	for i := 1; i < length; i++ {
		newPalette[i] = semiRandColor(newPalette[i-1].(color.RGBA), diff)
	}
	return newPalette
}

func main() {
	colorsAmount := 64
	maxDifference := 64 // maximum allowed difference between neighboring colors

	if len(os.Args) > 1 {
		colorsAmount, _ = strconv.Atoi(os.Args[1])
		if len(os.Args) > 2 {
			maxDifference, _ = strconv.Atoi(os.Args[2])
		}
	}

	palette = generatePalette(colorsAmount, maxDifference)
	file, err := os.Create("out.gif")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	lissajous(file)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(i % len(palette))
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to encode GIF: %v\n", err)
	}
}
