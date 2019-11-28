package main

import (
	"fmt"
	"github.com/smaTc/RayTracer/types"
)

func main() {
	fmt.Println("Hello")
	p := types.Vector{1, 1, 1, 1}
	fmt.Println(p.X)
	v := types.Tuple{1, 1, 1, 0}
	fmt.Println(v.W)
}
