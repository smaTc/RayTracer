package projectilegame

import (
	"fmt"
	"strconv"

	"github.com/smaTc/RayTracer/graphics"
)

func tick(e Environment, p Projectile) Projectile {
	pos := graphics.Add(p.Position, p.Velocity).(graphics.Point)
	vel := graphics.Add(graphics.Add(p.Velocity, e.Gravity), e.Wind).(graphics.Vector)
	return Projectile{Position: pos, Velocity: vel}
}

//Game Function
func Game() {
	vec := graphics.NewVector(1, 1, 0)
	proj := Projectile{Position: graphics.NewPoint(0, 1, 0), Velocity: graphics.NormVector(&vec)}

	env := Environment{Gravity: graphics.NewVector(0, -0.1, 0), Wind: graphics.NewVector(-0.01, 0, 0)}

	for i := 1; proj.Position.Y >= 0; i++ {
		proj = tick(env, proj)
		fmt.Println("Tick: "+strconv.Itoa(i)+" | Projectile Y-Val: ", proj.Position.Y)
	}
}
