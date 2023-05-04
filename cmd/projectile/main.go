package main

import (
	"flag"
	"os"

	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/image"
	"github.com/noahssarcastic/gort/pkg/math"
	"github.com/noahssarcastic/gort/pkg/ppm"
)

type projectile struct {
	position, velocity math.Tuple
}

type environment struct {
	gravity, wind math.Tuple
}

func tick(env environment, proj projectile) projectile {
	newPosition := proj.position.Add(proj.velocity)
	newVelocity := proj.velocity.Add(env.gravity).Add(env.wind)
	return projectile{newPosition, newVelocity}
}

var out = flag.String("o", "test.ppm", "output image path")

func main() {
	flag.Parse()

	proj := projectile{math.Point(0, 1, 0), math.Norm(math.Vector(1, 1, 0)).Mult(4)}
	env := environment{math.Vector(0, -0.1, 0), math.Vector(0, 0, 0)}
	img := image.New(200, 100)
	tickCount := 0
loop:
	for {
		tickCount += 1
		proj = tick(env, proj)
		x := int(proj.position.X())
		y := int(proj.position.Y())
		if x >= 0 && x < img.Width() && y >= 0 && y < img.Height() {
			img.Set(x, (img.Height()-1)-y, color.Red())
		}
		if proj.position.Y() <= 0 {
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
