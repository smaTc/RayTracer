package projectilegame

import (
	"fmt"

	"github.com/smaTc/RayTracer/graphicsmath"
)

func tick(e Environment, p Projectile) Projectile {
	pos := graphicsmath.Add(p.Position, p.Velocity).(graphicsmath.Point)
	vel := graphicsmath.Add(graphicsmath.Add(p.Velocity, e.Gravity), e.Wind).(graphicsmath.Vector)
	return Projectile{Position: pos, Velocity: vel}
}

//Game Function
func Game() {
	vec := graphicsmath.NewVector(1, 1, 0)
	proj := Projectile{Position: graphicsmath.NewPoint(0, 1, 0), Velocity: graphicsmath.NormVector(&vec)}

	env := Environment{Gravity: graphicsmath.NewVector(0, -0.1, 0), Wind: graphicsmath.NewVector(-0.01, 0, 0)}

	for i := 0; proj.Position.Y >= 0; i++ {
		proj = tick(env, proj)
		fmt.Println("Tick: "+string(i)+" | Projectile Y-Val: ", proj.Position.Y, 64)
	}
}
