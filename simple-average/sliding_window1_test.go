package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
