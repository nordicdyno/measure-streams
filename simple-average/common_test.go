package main

// common vars and initializaton for all tests

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
