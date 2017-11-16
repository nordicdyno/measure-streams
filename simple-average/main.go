package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var (
	filename   = flag.String("filename", "", "file with measurments")
	windowsize = flag.Int("window", 100, "window size")
)

func main() {
	flag.Parse()
	// var err error
	var r io.Reader
	if *filename == "" {
		r = os.Stdin
	} else {
		file, err := os.Open(*filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		r = file
	}

	wa := naiveWindowedAvg{windowSize: *windowsize}
	scanAndPrintResults(&wa, r, os.Stdout)
}

func scanAndPrintResults(wa *naiveWindowedAvg, r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	for i := 1; scanner.Scan(); i++ {
		measure := scanner.Text()
		f, err := strconv.ParseFloat(measure, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parse float line %v failed:\n", i)
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		// fmt.Fprintln(w, "Line", i)
		// fmt.Fprintf(w, "parsed: %v\n", f)
		wa.measure(f)
		fmt.Fprintf(w, "%v\n", wa.getMedian())
		// fmt.Fprintf(w, "median> %v\n", dw.getMedian())
	}
}
