package main

import (
	"sort"
)

type delaysWindow struct {
	measurements []float64
	size         int
	pointer      int

	// sort cache (to avoid rellocaton)
	sorted []float64
}

func (dw *delaysWindow) values() []float64 {
	window := make([]float64, len(dw.measurements))
	idx := 0
	for i := dw.pointer; i < len(dw.measurements); i++ {
		window[idx] = dw.measurements[i]
		idx++
	}
	for i := 0; i < dw.pointer; i++ {
		window[idx] = dw.measurements[i]
		idx++
	}
	return window
}

func (dw *delaysWindow) measures(values []float64) {
	for _, v := range values {
		dw.measure(v)
	}
}

func (dw *delaysWindow) measure(value float64) {
	if dw.measurements == nil {
		dw.measurements = make([]float64, 0, dw.size)
	}

	if dw.pointer == cap(dw.measurements) {
		dw.pointer = 0
	}
	if len(dw.measurements) < cap(dw.measurements) {
		dw.measurements = append(dw.measurements, value)
	} else {
		dw.measurements[dw.pointer] = value
	}
	dw.pointer++
}

func (dw *delaysWindow) getMedian() float64 {
	n := len(dw.measurements)
	if n < 2 {
		return -1
	}

	if dw.sorted == nil {
		dw.sorted = make([]float64, dw.size)
	}
	dw.sorted = dw.sorted[:n]
	copy(dw.sorted, dw.measurements)
	sort.Float64s(dw.sorted)
	if n%2 == 1 {
		return dw.sorted[n/2]
	}
	return (dw.sorted[n/2-1] + dw.sorted[n/2]) / 2
}
