package main

import (
	"github.com/noahssarcastic/tddraytracer/canvas"
	"github.com/noahssarcastic/tddraytracer/color"
	"github.com/noahssarcastic/tddraytracer/ppm"
	"github.com/noahssarcastic/tddraytracer/tuple"
)

type projectile struct {
	position, velocity tuple.Tuple
}

type environment struct {
	gravity, wind tuple.Tuple
}

func tick(env environment, proj projectile) projectile {
	newPosition := proj.position.Add(proj.velocity)
	newVelocity := proj.velocity.Add(env.gravity).Add(env.wind)
	return projectile{newPosition, newVelocity}
}

func main() {
	proj := projectile{tuple.Point(0, 1, 0), tuple.Normalize(tuple.Vector(1, 1, 0)).Multiply(4)}
	env := environment{tuple.Vector(0, -0.1, 0), tuple.Vector(0, 0, 0)}

	canv := canvas.New(200, 100)

	tickCount := 0
loop:
	for {
		tickCount += 1
		proj = tick(env, proj)

		// fmt.Printf("Tick: %v; Position: %v\n", tickCount, proj.position)

		x := int(proj.position.X())
		y := int(proj.position.Y())
		if x >= 0 && x < canv.Width() && y >= 0 && y < canv.Height() {
			println(x, y)
			canv.SetPixel(x, (canv.Height()-1)-y, color.Red())

		}

		if proj.position.Y() <= 0 {
			break loop
		}
	}

	ppm.CanvasToPixmap(canv).WritePPM()
}
