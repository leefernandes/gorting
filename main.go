package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/leefernandes/gorting/pkg/bubble"
	"github.com/leefernandes/gorting/pkg/custom3"
	"github.com/leefernandes/gorting/pkg/quick"
	"github.com/leefernandes/gorting/pkg/util"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Sorter = func([]int) []int

func main() {
	n := 1
	size := 10000001

	//smpl := []int{8,9,2,-5,1,0,1,7,11,2,5,9,2,-3,1,9}
	// debug(smpl, bubbleSort)
	// return

	for i := 0; i < n; i++ {
		runSorts(size)
	}
}

func runDebug() {
	smpl := []int{8, 9, 2, -5, 1, 0, 1, 7, 11, 2, 5, 9, 2, -3, 1, 9}
	debug(smpl, bubble.Sort)
	return
}

func runSorts(size int) {
	og := util.GenerateRandomSlice(size)

	//bubbleSort(og)
	//return

	p := message.NewPrinter(language.English)
	p.Printf("sort []int of length %d\n", size)

	stdSorted := runSort(og, stdSort)

	//defer compareSorted("bubble.Sort", runSort(og, bubble.Sort), stdSorted, og)
	defer compareSorted("quick.Sort", runSort(og, quick.Sort), stdSorted, og)
	// defer compareSorted("custom1.Sort", runSort(og, custom1.Sort), stdSorted, og)
	// defer compareSorted("custom1.Sortb", runSort(og, custom1.Sortb), stdSorted, og)
	// defer compareSorted("custom2.Sort", runSort(og, custom2.Sort), stdSorted, og)
	defer compareSorted("custom3.Sort", runSort(og, custom3.Sort), stdSorted, og)
}

func stdSort(s []int) []int {
	start := time.Now()
	sort.Ints(s)
	elapsed := time.Since(start)
	fmt.Printf("sort.Ints %s\n", elapsed)
	return s
}

func runSort(s []int, sorter Sorter) []int {
	return sorter(util.CopySlice(s))
}

func compareSorted(name string, customSorted []int, stdSorted []int, original []int) {
	if !reflect.DeepEqual(customSorted, stdSorted) {
		fmt.Println(name, "fail", unmatchedIndices(customSorted, stdSorted))
		fmt.Println("")
		fmt.Println(name, len(customSorted), customSorted)
		fmt.Println("")
		fmt.Println("stdSorted:", len(stdSorted), stdSorted)
		fmt.Println("")

		list := "[]int{"
		for _, v := range original {
			list += fmt.Sprintf("%v,", v)
		}
		list += "}"
		fmt.Println("debug:", list)
	} else {
		fmt.Println(name, "ok")
	}
}

func unmatchedIndices(a, b []int) []string {
	indices := []string{}
	for i := range a {
		if a[i] != b[i] {
			indices = append(indices, fmt.Sprintf("%v got %v should be %v", i, a[i], b[i]))
		}
	}
	return indices
}

func debug(s []int, sorter Sorter) {
	s1 := util.CopySlice(s)
	s2 := util.CopySlice(s)

	stdSorted := stdSort(s1)
	customSorted := sorter(s2)
	fmt.Println(reflect.DeepEqual(stdSorted, customSorted), len(stdSorted), len(customSorted))
	fmt.Println("original slice:")
	fmt.Println(s)
	fmt.Println("")
	fmt.Println("default sorted:")
	fmt.Println(stdSorted)
	fmt.Println("")
	fmt.Println("custom sorted:")
	fmt.Println(customSorted)
	fmt.Println("")

	for i, v := range stdSorted {
		if customSorted[i] != v {
			fmt.Println("oh shit", i, v, customSorted[i])
			break
		}
	}
}
