package main

import (
	"fmt"
	"github.com/smaTc/RayTracer/types"
)

func main() {
	fmt.Println("Hello")
	p := types.Vector{1, 1, 1, 1}
	fmt.Println(p.X)
	v := types.NewTuple{1, 1, 1, 0}.(types.Vector)
	fmt.Println(v.W)
}
