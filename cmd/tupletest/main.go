package main

import (
	"fmt"

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
	proj := projectile{tuple.Point(0, 1, 0), tuple.Normalize(tuple.Vector(1, 1, 0))}
	env := environment{tuple.Vector(0, -0.1, 0), tuple.Vector(-0.01, 0, 0)}

	tickCount := 0
loop:
	for {
		tickCount += 1
		proj = tick(env, proj)
		fmt.Printf("Tick: %v; Position: %v\n", tickCount, proj.position)
		if proj.position.Y() <= 0 {
			break loop
		}
	}
}
