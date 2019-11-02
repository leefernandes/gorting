package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func copySlice(s1 []int) []int {
	s2 := make([]int, len(s1))
	copy(s2, s1)
	return s2
}

func generateRandomSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(10000001) // - rand.Intn(10000001)
	}
	return slice
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
	s1 := copySlice(s)
	s2 := copySlice(s)

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
