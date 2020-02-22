package projectilegame

import (
	"fmt"
	"strconv"

	"github.com/smaTc/RayTracer/graphics"
)

func tick(e Environment, p *Projectile) {
	pos := graphics.Add(p.Position, p.Velocity).(graphics.Point)
	vel := graphics.Add(graphics.Add(p.Velocity, e.Gravity), e.Wind).(graphics.Vector)
	p.Position = pos
	p.Velocity = vel
	//return Projectile{Position: pos, Velocity: vel}
}

//Game Function
func Game() {
	vec := graphics.NewVector(1, 1, 0)
	proj := Projectile{Position: graphics.NewPoint(0, 1, 0), Velocity: *vec.Normalize()}

	env := Environment{Gravity: graphics.NewVector(0, -0.02, 0), Wind: graphics.NewVector(0.001, 0, 0)}

	c := graphics.CreateCanvas(200, 200)

	for i := 1; proj.Position.Y >= 0.0; i++ {
		fmt.Println("Tick: "+strconv.Itoa(i)+" | Projectile Y-Val: ", proj.Position.Y)
		c.WritePixelToCanvas(int(proj.Position.X), int(proj.Position.Y), &graphics.Color{Red: 1})
		tick(env, &proj)
	}
	graphics.SaveCanvasAsPPM(&c, "/tmp/", "test")
}
