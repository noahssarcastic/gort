package main

import (
	"flag"
	"os"

	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/image"
	"github.com/noahssarcastic/gort/pkg/ppm"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

type projectile struct {
	position, velocity tuple.Tuple
}

type environment struct {
	gravity, wind tuple.Tuple
}

func tick(env environment, proj projectile) projectile {
	newPosition := tuple.Add(proj.position, proj.velocity)
	newVelocity := tuple.Add(
		tuple.Add(proj.velocity, env.gravity),
		env.wind)
	return projectile{newPosition, newVelocity}
}

var out = flag.String("o", "test.ppm", "output image path")

func main() {
	flag.Parse()

	proj := projectile{tuple.Point(0, 1, 0), tuple.Mult(tuple.Norm(tuple.Vector(1, 1, 0)), 4)}
	env := environment{tuple.Vector(0, -0.1, 0), tuple.Vector(0, 0, 0)}
	img := image.New(200, 100)
	tickCount := 0
loop:
	for {
		tickCount += 1
		proj = tick(env, proj)
		x := int(proj.position.X)
		y := int(proj.position.Y)
		if x >= 0 && x < img.Width() && y >= 0 && y < img.Height() {
			img.Set(x, (img.Height()-1)-y, color.Red)
		}
		if proj.position.Y <= 0 {
			break loop
		}
	}
	pm := image.ImageToPixelMap(*img)

	f, err := os.OpenFile(*out,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0755)
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
