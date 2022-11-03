package cmd

import (
	"github.com/noahssarcastic/gort/image"
	"github.com/noahssarcastic/gort/math"
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

func RunProjectile() {
	proj := projectile{math.Point(0, 1, 0), math.Norm(math.Vector(1, 1, 0)).Mult(4)}
	env := environment{math.Vector(0, -0.1, 0), math.Vector(0, 0, 0)}

	canv := image.NewCanvas(200, 100)

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
			canv.SetPixel(x, (canv.Height()-1)-y, image.Red())

		}

		if proj.position.Y() <= 0 {
			break loop
		}
	}

	image.CanvasToPixmap(canv).WritePPM()
}
