package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nordicdyno/measure-streams/modeling"
)

func TestSortedDebug1(t *testing.T) {
	wa := sortedWindowedAvg{
		windowSize: 5,
	}
	var debugData = []struct {
		value        float64
		sortedExpect []float64
	}{
		{200, []float64{200}},
		{99, []float64{99, 200}},
		{101, []float64{99, 101, 200}},
		{101, []float64{99, 101, 101, 200}},
		{120, []float64{99, 101, 101, 120, 200}},
		{101, []float64{99, 101, 101, 101, 120}},
		// {120, []float64{101, 110, 120}},
		// {115, []float64{110, 120, 115}},
	}
	for step, c := range debugData {
		fmt.Printf("TestSortedDebug1: ================  step %v ===============\n", step)
		wa.measure(c.value)
		// fmt.Printf("measurements => %v\n", wa.measurements)

		// values := wa.values()
		assert.Equal(t, c.sortedExpect, wa.sorted, fmt.Sprintf("step N %v", step))
	}
}

func TestSortedDebug2(t *testing.T) {
	var debugData = []struct {
		name     string
		initial  []float64
		newvalue float64
	}{
		// {
		// 	name:     "rmindex==newindex",
		// 	initial:  []float64{100, 50, 99, 120, 222},
		// 	newvalue: 150,
		// },
		// {
		// 	name:     "add on the right: rmindex < newindex",
		// 	initial:  []float64{100, 50, 99, 120, 222},
		// 	newvalue: 500,
		// },
		{
			name:     "Special case: newindex-1 == rmindex",
			initial:  []float64{120, 50, 99, 100, 222},
			newvalue: 200,
		},
		{
			name:     "add on the right: rmindex > newindex",
			initial:  []float64{100, 90, 80, 120, 50, 99, 100, 222, 300, 400, 1000},
			newvalue: 200,
		},
	}
	for step, dat := range debugData {
		fmt.Printf("TestSortedDebug2: ================  step %v ===============\n", step)
		wa := sortedWindowedAvg{
			windowSize: len(dat.initial),
		}
		wa.init(dat.initial)
		// if wa.sorted == nil {
		// wa.sorted = make([]float64, wa.windowSize)
		// copy(wa.sorted, wa.measurements)
		// sort.Float64s(wa.sorted)

		expect := make([]float64, wa.windowSize)
		copy(expect, dat.initial)
		expect[0] = dat.newvalue
		sort.Float64s(expect)

		// fmt.Printf("%v\n", wa)
		wa.measure(dat.newvalue)
		// 	wa.measure(c.value)
		// fmt.Printf("measurements => %v\n", wa.measurements)
		// fmt.Printf("sorted => %v\n", wa.sorted)

		// 	// values := wa.values()
		assert.Equal(t, expect, wa.sorted, dat.name)
	}
}

func TestSortedInitialState(t *testing.T) {
	var wa sortedWindowedAvg
	if len(wa.values()) > 0 {
		t.Error("initial values failed")
	}
}

func TestSortedTest1(t *testing.T) {
	wa := sortedWindowedAvg{windowSize: 3}
	for step, c := range measurecases {
		wa.measure(c.value)
		values := wa.values()
		assert.Equal(t, c.expect, values, fmt.Sprintf("step N %v", step))
	}
}

func TestSortedMedian(t *testing.T) {
	for step, c := range mediancases {
		wa := sortedWindowedAvg{windowSize: 1000}
		wa.measures(c.values)
		assert.Equal(t, c.median, wa.getMedian(), fmt.Sprintf("step N %v", step))
	}
}

// * Sliding window size for Test 2 is 100
// * Sliding window size for Test 3 is 1000
// * Sliding window size for Test 4 is 10000

var sortedResult float64

func benchSortedWindow(b *testing.B, size int) {
	b.StopTimer()
	data, err := modeling.GenParetoN(0.02, 0.98, 100*1000)
	if err != nil {
		b.Error(err)
	}

	wa := sortedWindowedAvg{windowSize: size}

	for _, elem := range data {
		wa.measure(elem)
	}

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		sortedResult = wa.getMedian()
	}
}

func BenchmarkSortedWindow100(b *testing.B) {
	benchSortedWindow(b, 100)
}

func BenchmarkSortedWindow1k(b *testing.B) {
	benchSortedWindow(b, 1000)
}

func BenchmarkSortedWindow10k(b *testing.B) {
	benchSortedWindow(b, 10000)
}
