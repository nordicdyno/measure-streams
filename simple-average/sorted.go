package main

import (
	"log"
	"sort"
)

type sortedWindowedAvg struct {
	measurements []float64
	windowSize   int
	pointer      int

	// TODO: withCache bool
	// sort cache (to avoid rellocaton)
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

func findindex(a []float64, x float64) int {
	i := sort.SearchFloat64s(a, x)
	// not sure here is this possible to find value
	if i > 0 {
		if a[i-1] == x {
			log.Println("found the same value on the left")
		}
	}
	// it's ok, we expect this
	// if i+1 < len(a) {
	// 	if a[i+1] == x {
	// 		log.Println("found the same value on the right")
	// 	}
	// }
	return i
}

func (wa *sortedWindowedAvg) addSorted(value float64) {
	// fmt.Println("      . . . . . . . . . . .")
	// fmt.Printf("START OF addSorted(%v) -> sorted: %+v\n", value, wa.sorted)

	// copy(wa.sorted, wa.measurements)
	// sort.Float64s(wa.sorted)

	if len(wa.sorted) < cap(wa.sorted) {
		// 1) find index for new value
		// 2) append sorted array with new value
		// 3) if found index is not last shift all index:end-1 to right inde
		//    and put new value ins place
		// fmt.Println(" > addSorted: branch1")
		newindex := sort.SearchFloat64s(wa.sorted, value)
		// fmt.Printf(" > value=%v newindex=%v\n", value, newindex)
		was := len(wa.sorted)
		wa.sorted = append(wa.sorted, value)
		if newindex != was {
			// fmt.Printf("shift [%v:%v] -> [%v:%v]\n", newindex+1,len(wa.sorted), newindex,was)
			copy(wa.sorted[newindex+1:len(wa.sorted)], wa.sorted[newindex:was])
			wa.sorted[newindex] = value
		}
		return
	}

	// fmt.Println(" > addSorted: branch2")
	rmvalue := wa.measurements[wa.pointer]
	// fmt.Printf(" > rmvalue=%v; value=%v\n", rmvalue, value)
	if rmvalue == value {
		return
	}

	rmindex := sort.SearchFloat64s(wa.sorted, rmvalue)
	newindex := sort.SearchFloat64s(wa.sorted, value)
	// fmt.Printf(" > rmindex=%v; newindex=%v\n", rmindex, newindex)
	if newindex == rmindex {
		wa.sorted[rmindex] = value
		return
	}

	if rmindex < newindex {
		// just save
		if newindex-1 == rmindex {
			wa.sorted[rmindex] = value
			return
		}
		// if rmindex < newindex -> shift elements from [rmindex+1:newindex-1] to [rmindex:newindex-2]
		// fmt.Printf("shift [%v:%v) -> [%v:%v)\n", rmindex+1, newindex, rmindex, newindex-1)
		copy(wa.sorted[rmindex:newindex-1], wa.sorted[rmindex+1:newindex])
		wa.sorted[newindex-1] = value
		return
	}
	// rmindex > newindex
	// fmt.Printf("shift [%v:%v] -> [%v:%v]\n", newindex, rmindex, newindex+1, rmindex+1)
	copy(wa.sorted[newindex+1:rmindex+1], wa.sorted[newindex:rmindex])
	wa.sorted[newindex] = value
	// 0) figure out of removed value
	// 1) if value == removed value, nothing to do
	// 2) find rmindex of removed value
	// 3) find newindex for new value
	// 4.a) if rmindex == newindex just store value on index1
	// 4.b) if rmindex < newindex shift elements from [rmindex+1:newindex-1] to [rmindex:newindex-2]
	//      and store value in index2-1
	// 4c) if rmindex > index2 shift elements from [newindex:rmindex-1] to [newindex+1:rmindex]
	//
	//    and place new value in place
	// move to defer if want do debug
	// fmt.Printf("END OF addSorted(%v) -> sorted: %+v\n", value, wa.sorted)
	// fmt.Println("      . . . . . . . . . . .")
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

	// wa.add
	wa.pointer++
}

// useful for testing purposes
func (wa *sortedWindowedAvg) init(values []float64) {
	length := len(values)
	// fmt.Println("length:", length)
	if length > wa.windowSize {
		length = wa.windowSize
	}
	// fmt.Println("length:", length)
	wa.measurements = make([]float64, length, wa.windowSize)
	wa.sorted = make([]float64, length, wa.windowSize)

	// if len(values) > wa.windowSize {
	// fmt.Printf("copy [%v:%v]\n", len(values)-wa.windowSize, len(values))
	copy(wa.measurements, values[len(values)-wa.windowSize:len(values)])
	copy(wa.sorted, wa.measurements)
	sort.Float64s(wa.sorted)

	// }
	wa.pointer = cap(wa.measurements)
	// fmt.Printf("%v => %v\n", values, wa.measurements)
}

func (wa *sortedWindowedAvg) getMedian() float64 {
	n := len(wa.measurements)
	if n < 2 {
		return -1
	}

	// if wa.sorted == nil {
	// 	wa.sorted = make([]float64, wa.windowSize)
	// }
	// wa.sorted = wa.sorted[:n]
	// copy(wa.sorted, wa.measurements)
	// sort.Float64s(wa.sorted)
	if n%2 == 1 {
		return wa.sorted[n/2]
	}
	return (wa.sorted[n/2-1] + wa.sorted[n/2]) / 2
}
