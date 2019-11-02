package quick

import (
	"fmt"
	"time"
)

func Sort(s []int) []int {
	start := time.Now()

	Quicksort(s, 0, len(s))

	elapsed := time.Since(start)
	fmt.Printf("quick.Sort %s\n", elapsed)
	//fmt.Println(len(s), s)
	return s
}

func Quicksort(s []int, start int, length int) {
	end := length - 1
	if start >= end {
		return
	}
	p := Partition(s, start, length)
	Quicksort(s, start, p)
	Quicksort(s, p+1, length)
}

func Partition(s []int, start int, length int) int {
	end := length - 1
	p := s[end]
	pI := start
	for i := start; i < end; i++ {
		v := s[i]
		if v <= p {
			s[i], s[pI] = s[pI], s[i]
			pI++
		}
	}
	s[end], s[pI] = s[pI], s[end]
	return pI
}
