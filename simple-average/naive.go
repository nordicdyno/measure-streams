package main

import (
	"sort"
)

type naiveWindowedAvg struct {
	measurements []float64
	windowSize   int
	pointer      int

	// TODO: withCache bool
	// sort cache (to avoid rellocaton)
	sorted []float64
}

func (wa *naiveWindowedAvg) values() []float64 {
	window := make([]float64, len(wa.measurements))
	idx := 0
	for i := wa.pointer; i < len(wa.measurements); i++ {
		window[idx] = wa.measurements[i]
		idx++
	}
	for i := 0; i < wa.pointer; i++ {
		window[idx] = wa.measurements[i]
		idx++
	}
	return window
}

func (wa *naiveWindowedAvg) measures(values []float64) {
	for _, v := range values {
		wa.measure(v)
	}
}

func (wa *naiveWindowedAvg) measure(value float64) {
	if wa.measurements == nil {
		wa.measurements = make([]float64, 0, wa.windowSize)
	}

	if wa.pointer == cap(wa.measurements) {
		wa.pointer = 0
	}
	if len(wa.measurements) < cap(wa.measurements) {
		wa.measurements = append(wa.measurements, value)
	} else {
		wa.measurements[wa.pointer] = value
	}
	wa.pointer++
}

func (wa *naiveWindowedAvg) getMedian() float64 {
	n := len(wa.measurements)
	if n < 2 {
		return -1
	}

	if wa.sorted == nil {
		wa.sorted = make([]float64, wa.windowSize)
	}
	wa.sorted = wa.sorted[:n]
	copy(wa.sorted, wa.measurements)
	sort.Float64s(wa.sorted)
	if n%2 == 1 {
		return wa.sorted[n/2]
	}
	return (wa.sorted[n/2-1] + wa.sorted[n/2]) / 2
}
