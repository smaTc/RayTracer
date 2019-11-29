package projectilegame

import "github.com/smaTc/RayTracer/graphics"

//Environment struct
type Environment struct {
	Gravity, Wind graphics.Vector
}

//Projectile struct
type Projectile struct {
	Position graphics.Point
	Velocity graphics.Vector
}
