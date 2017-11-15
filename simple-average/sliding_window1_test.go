package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nordicdyno/measure-streams/modeling"
)

func TestInitialState(t *testing.T) {
	var dw delaysWindow
	if len(dw.values()) > 0 {
		t.Error("initial values failed")
	}
}

var measurecases = []struct {
	value  float64
	expect []float64
}{
	{100, []float64{100}},
	{102, []float64{100, 102}},
	// {101, []float64{100, 102, 101}},
	// {110, []float64{102, 101, 110}},
	// {120, []float64{101, 110, 120}},
	// {115, []float64{110, 120, 115}},
}

func TestTest1(t *testing.T) {
	dw := delaysWindow{size: 3}
	for step, c := range measurecases {
		dw.measure(c.value)
		values := dw.values()
		assert.Equal(t, c.expect, values, fmt.Sprintf("step N %v", step))
	}
}

var mediancases = []struct {
	values []float64
	median float64
}{
	{[]float64{}, -1},
	{[]float64{100}, -1},
	{[]float64{100, 102}, 101},
	{[]float64{200, 100, 102, 100}, 101},
	{[]float64{102, 101, 110}, 102},
	{[]float64{99, 110, 115, 118, 120, 345, 567, 100500}, 119},
}

func TestTestMedian(t *testing.T) {
	for step, c := range mediancases {
		dw := delaysWindow{size: 1000}
		dw.measures(c.values)
		assert.Equal(t, c.median, dw.getMedian(), fmt.Sprintf("step N %v", step))
	}
}

// * Sliding window size for Test 2 is 100
// * Sliding window size for Test 3 is 1000
// * Sliding window size for Test 4 is 10000

var result float64

func benchWindow(b *testing.B, size int) {
	b.StopTimer()
	data, err := modeling.GenParetoN(0.02, 0.98, 100*1000)
	if err != nil {
		b.Error(err)
	}

	dw := delaysWindow{size: size}

	for _, elem := range data {
		dw.measure(elem)
	}

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		result = dw.getMedian()
	}
}

func BenchmarkWindow100(b *testing.B) {
	benchWindow(b, 100)
}

func BenchmarkWindow1k(b *testing.B) {
	benchWindow(b, 1000)
}

func BenchmarkWindow10k(b *testing.B) {
	benchWindow(b, 10000)
}
