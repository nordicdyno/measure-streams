package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nordicdyno/measure-streams/modeling"
)

func TestNaiveInitialState(t *testing.T) {
	var dw delaysWindow
	if len(dw.values()) > 0 {
		t.Error("initial values failed")
	}
}

func TestNaiveTest1(t *testing.T) {
	dw := delaysWindow{size: 3}
	for step, c := range measurecases {
		dw.measure(c.value)
		values := dw.values()
		assert.Equal(t, c.expect, values, fmt.Sprintf("step N %v", step))
	}
}

func TestNaiveTestMedian(t *testing.T) {
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

func benchNaiveWindow(b *testing.B, size int) {
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

func BenchmarkNaiveWindow100(b *testing.B) {
	benchNaiveWindow(b, 100)
}

func BenchmarkNaiveWindow1k(b *testing.B) {
	benchNaiveWindow(b, 1000)
}

func BenchmarkNaiveWindow10k(b *testing.B) {
	benchNaiveWindow(b, 10000)
}
