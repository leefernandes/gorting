package main

import (
	"log"
	"time"
)

func quicksort(s []int) []int {
	start := time.Now()

	quick(s, 0, len(s))

	elapsed := time.Since(start)
	log.Printf("quicksort took %s\n", elapsed)
	//fmt.Println(len(s), s)
	return s
}

func quick(s []int, start int, length int) {
	end := length - 1
	if start >= end {
		return
	}
	p := partition(s, start, length)
	quick(s, start, p)
	quick(s, p+1, length)
}

func partition(s []int, start int, length int) int {
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
