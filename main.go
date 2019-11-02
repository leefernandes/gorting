package main

import (
	"log"
	"sort"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Sorter = func([]int) []int

func main() {
	n := 1
	size := 100001

	//smpl := []int{8,9,2,-5,1,0,1,7,11,2,5,9,2,-3,1,9}
	// debug(smpl, bubbleSort)
	// return

	for i := 0; i < n; i++ {
		runSorts(size)
	}
}

func runDebug() {
	smpl := []int{8, 9, 2, -5, 1, 0, 1, 7, 11, 2, 5, 9, 2, -3, 1, 9}
	debug(smpl, bubbleSort)
	return
}

func runSorts(size int) {

	og := generateRandomSlice(size)

	//bubbleSort(og)
	//return

	s1 := copySlice(og)
	s2 := copySlice(og)
	s3 := copySlice(og)

	p := message.NewPrinter(language.English)
	p.Printf("sort []int of length %d\n", size)

	stdSorted := stdSort(s1)
	quickSorted := quicksort(s3)
	bubbleSorted := bubbleSort(s2)

	compareSorted("quick sort", quickSorted, stdSorted, og)
	compareSorted("bubble sort", bubbleSorted, stdSorted, og)
}

func stdSort(s []int) []int {
	start := time.Now()

	sort.Ints(s)

	elapsed := time.Since(start)
	log.Printf("sort.Ints sort took %s\n", elapsed)
	return s
}
