package modeling

import "github.com/atgjack/prob"

func GenParetoN(scale float64, shape float64, n int) ([]float64, error) {
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
