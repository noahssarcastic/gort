package main

import (
	"flag"
	"os"

	"github.com/noahssarcastic/gort/pkg/ppm"
)

var out = flag.String("o", "test.ppm", "output image path")

func main() {
	flag.Parse()

	w, h := 100, 100
	pm := ppm.New(w, h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b :=
				int((float64(x)/float64(w))*ppm.MaxColor),
				0,
				int((float64(y)/float64(h))*ppm.MaxColor)
			pm.Set(x, y, r, g, b)
		}
	}

	f, err := os.OpenFile(*out, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()
	ppm.WritePPM(f, pm)
}
