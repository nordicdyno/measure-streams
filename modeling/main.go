package main

import (
	"flag"
	"fmt"

	"github.com/atgjack/prob"
)

var (
	scale = flag.Float64("scale", 0.02, "mean")
	shape = flag.Float64("shape", 0.98, "shape")
	count = flag.Int("c", 100, "iterations")
)

func main() {
	flag.Parse()
	elems, err := genN(*scale, *shape, *count)
	if err != nil {
		panic(err)
	}
	for _, e := range elems {
		fmt.Printf("%.04f\n", e)
	}
}

func genN(scale float64, shape float64, n int) ([]float64, error) {
	var elems []float64
	dist, err := prob.NewPareto(scale, shape)
	if err != nil {
		return nil, err
	}
	for i := 0; i < n; i++ {
		elems = append(elems, dist.Random())
		// fmt.Printf("%.2f\n", poisson.Random())
	}
	return elems, nil
}
