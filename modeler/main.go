package main

import (
	"flag"
	"fmt"

	"github.com/nordicdyno/measure-streams/modeling"
)

var (
	scale = flag.Float64("scale", 0.02, "mean")
	shape = flag.Float64("shape", 0.98, "shape")
	count = flag.Int("c", 100, "iterations")
)

func main() {
	flag.Parse()
	elems, err := modeling.GenParetoN(*scale, *shape, *count)
	if err != nil {
		panic(err)
	}
	for _, e := range elems {
		fmt.Printf("%.04f\n", e)
	}
}
