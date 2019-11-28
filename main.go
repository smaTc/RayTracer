package main

import (
	"fmt"

	"github.com/smaTc/RayTracer/graphics"
	"github.com/smaTc/RayTracer/types"
)

func main() {
	fmt.Println("Hello")
	v, err := graphics.Tuple(1, 1, 1, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(v.(types.Vector).W)

}
