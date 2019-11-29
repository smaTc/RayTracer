package projectilegame

import "github.com/smaTc/RayTracer/graphicsmath"

//Environment struct
type Environment struct {
	Gravity, Wind graphicsmath.Vector
}

//Projectile struct
type Projectile struct {
	Position graphicsmath.Point
	Velocity graphicsmath.Vector
}
