package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var pallete = []color.Color{
	color.White,
	color.RGBA{12, 177, 205, 255},
	color.RGBA{255, 108, 0, 255},
	color.RGBA{255, 209, 1, 255},
}

const (
	whiteIndex     = 0
	acquaBlueIndex = 1
	orangeIndex    = 2
	yelloyIndex    = 3
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete revolutions of x oscilator
		res     = 0.001 // angular resolution
		size    = 100   // image canvas [-size..+size]
		nframes = 32    // number of animation frames
		delay   = 12    // time between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // y oscilator relative frequency
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(1 + i%3)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // intentionally ignoring errors
}
