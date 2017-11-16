package main

import (
	"sort"
)

type sortedWindowedAvg struct {
	measurements []float64
	windowSize   int
	pointer      int

	// sorted data to avoid sort on demand complexity
	sorted []float64
}

func (wa *sortedWindowedAvg) values() []float64 {
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

func (wa *sortedWindowedAvg) measures(values []float64) {
	for _, v := range values {
		wa.measure(v)
	}
}

func (wa *sortedWindowedAvg) addSorted(value float64) {
	if len(wa.sorted) < cap(wa.sorted) {
		// in case if buffer not full, adds new value:
		// 1) find index for new value
		// 2) append sorted array with new value
		// 3) if found index is not last shift all index:end-1 to right  and put new value in place
		newindex := sort.SearchFloat64s(wa.sorted, value)
		was := len(wa.sorted)
		wa.sorted = append(wa.sorted, value)
		if newindex != was {
			copy(wa.sorted[newindex+1:len(wa.sorted)], wa.sorted[newindex:was])
			wa.sorted[newindex] = value
		}
		return
	}

	// in case if buffer full, removes old and adds new value
	// 0) figure out which value to remove
	// 1) if value == removed value, END
	// 2) find rmindex of removed value and newindex for new value
	// 3) if newindex == rmindex: store value on rmindex, END
	rmvalue := wa.measurements[wa.pointer]
	if rmvalue == value {
		return
	}
	rmindex := sort.SearchFloat64s(wa.sorted, rmvalue)
	newindex := sort.SearchFloat64s(wa.sorted, value)
	if newindex == rmindex {
		wa.sorted[rmindex] = value
		return
	}

	// 4) if rmindex < newindex
	// 4.a) check special case newindex-1 == rmindex, if true save value with rmindex, END
	// 4.b) shift elements in dapasone [rmindex+1:newindex-1] to left,
	//      store value in newindex-1, END
	if rmindex < newindex {
		if newindex-1 == rmindex {
			wa.sorted[rmindex] = value
			return
		}
		copy(wa.sorted[rmindex:newindex-1], wa.sorted[rmindex+1:newindex])
		wa.sorted[newindex-1] = value
		return
	}

	// 5) if rmindex > index2
	// shift elements from diapason [newindex:rmindex-1] to right,
	// save new value in newindex
	copy(wa.sorted[newindex+1:rmindex+1], wa.sorted[newindex:rmindex])
	wa.sorted[newindex] = value
}

func (wa *sortedWindowedAvg) measure(value float64) {
	if wa.measurements == nil {
		wa.measurements = make([]float64, 0, wa.windowSize)
		wa.sorted = make([]float64, 0, wa.windowSize)
	}

	if wa.pointer == cap(wa.measurements) {
		wa.pointer = 0
	}
	wa.addSorted(value)

	if len(wa.measurements) < cap(wa.measurements) {
		wa.measurements = append(wa.measurements, value)
	} else {
		wa.measurements[wa.pointer] = value
	}

	wa.pointer++
}

// useful for testing purposes
func (wa *sortedWindowedAvg) init(values []float64) {
	length := len(values)
	if length > wa.windowSize {
		length = wa.windowSize
	}
	wa.measurements = make([]float64, length, wa.windowSize)
	wa.sorted = make([]float64, length, wa.windowSize)

	copy(wa.measurements, values[len(values)-wa.windowSize:len(values)])
	copy(wa.sorted, wa.measurements)
	sort.Float64s(wa.sorted)

	wa.pointer = cap(wa.measurements)
}

func (wa *sortedWindowedAvg) getMedian() float64 {
	n := len(wa.measurements)
	if n < 2 {
		return -1
	}

	if n%2 == 1 {
		return wa.sorted[n/2]
	}
	return (wa.sorted[n/2-1] + wa.sorted[n/2]) / 2
}
