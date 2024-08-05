// Exercise 1.12: Modify the Lissajous server to read parameter values from the URL. For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets the
//number of cycles to 20 instead of the default 5. Use the strconv.Atoi function to convert the
//string parameter into an integer. You can see its documentation with go doc strconv.Atoi.

// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var startColor = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}
var palette []color.Color
var (
	generationParameters = map[string]string{
		"cycles":  "5",     // number of complete x oscillator revolutions
		"res":     "0.001", // angular resolution
		"size":    "100",   // image canvas covers [-size..+size]
		"nframes": "64",    // number of animation frames
		"delay":   "8"}
)

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

func Lissajous(out io.Writer) {
	var (
		cycles, _  = strconv.Atoi(generationParameters["cycles"])        // number of complete x oscillator revolutions
		res, _     = strconv.ParseFloat(generationParameters["res"], 64) // angular resolution
		size, _    = strconv.Atoi(generationParameters["size"])          // image canvas covers [-size..+size]
		nframes, _ = strconv.Atoi(generationParameters["nframes"])       // number of animation frames
		delay, _   = strconv.Atoi(generationParameters["delay"])         // delay between frames in 10ms units
	)
	colorsAmount := 64
	maxDifference := 32 // maximum allowed difference between neighboring colors

	if len(os.Args) > 1 {
		colorsAmount, _ = strconv.Atoi(os.Args[1])
		if len(os.Args) > 2 {
			maxDifference, _ = strconv.Atoi(os.Args[2])
		}
	}
	palette = generatePalette(colorsAmount, maxDifference)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(i % len(palette))
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), colorIndex)
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

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for k := range generationParameters {
			if v := r.URL.Query().Get(k); v != "" {
				generationParameters[k] = v
			}
		}
		Lissajous(w)
	})
	// each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
