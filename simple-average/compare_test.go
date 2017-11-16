package main

import (
	"fmt"
	"testing"

	"github.com/nordicdyno/measure-streams/modeling"
	"github.com/stretchr/testify/assert"
)

func TestSideBySideMedian(t *testing.T) {
	data, err := modeling.GenParetoN(0.02, 0.98, 1000*1000)
	if err != nil {
		panic(err)
	}
	var size = int(60)
	sortedWA := sortedWindowedAvg{windowSize: size}
	naiveWA := naiveWindowedAvg{windowSize: size}

	var (
		step int
		elem float64
	)
	for step, elem = range data {
		naiveWA.measure(elem)
		sortedWA.measure(elem)
		assert.Equal(t, sortedWA.getMedian(), naiveWA.getMedian(), fmt.Sprintf("failed on step N %v", step))
	}
	fmt.Println("passed steps", step+1)
}
