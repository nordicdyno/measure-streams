package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nordicdyno/measure-streams/modeling"
)

func TestNaiveInitialState(t *testing.T) {
	var wa naiveWindowedAvg
	if len(wa.values()) > 0 {
		t.Error("initial values failed")
	}
}

func TestNaiveTest1(t *testing.T) {
	wa := naiveWindowedAvg{windowSize: 3}
	for step, c := range measurecases {
		wa.measure(c.value)
		values := wa.values()
		assert.Equal(t, c.expect, values, fmt.Sprintf("step N %v", step))
	}
}

func TestNaiveTestMedian(t *testing.T) {
	for step, c := range mediancases {
		wa := naiveWindowedAvg{windowSize: 1000}
		wa.measures(c.values)
		assert.Equal(t, c.median, wa.getMedian(), fmt.Sprintf("step N %v", step))
	}
}

// * Sliding window size for Test 2 is 100
// * Sliding window size for Test 3 is 1000
// * Sliding window size for Test 4 is 10000

var result float64

func benchNaiveWindow(b *testing.B, size int) {
	b.StopTimer()
	data, err := modeling.GenParetoN(0.02, 0.98, 100*1000)
	if err != nil {
		b.Error(err)
	}

	wa := naiveWindowedAvg{windowSize: size}

	for _, elem := range data {
		wa.measure(elem)
	}

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		result = wa.getMedian()
	}
}

func BenchmarkNaiveWindow100(b *testing.B) {
	benchNaiveWindow(b, 100)
}

func BenchmarkNaiveWindow1k(b *testing.B) {
	benchNaiveWindow(b, 1000)
}

func BenchmarkNaiveWindow10k(b *testing.B) {
	benchNaiveWindow(b, 10000)
}
