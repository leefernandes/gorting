package main

import (
	"log"
	"time"
)

func bubbleSort(s []int) []int {
	start := time.Now()

	for i := len(s); i > 0; i-- {
		for j := 1; j < i; j++ {
			if s[j-1] > s[j] {
				intermediate := s[j]
				s[j] = s[j-1]
				s[j-1] = intermediate
			}
		}
	}
	elapsed := time.Since(start)
	log.Printf("bubble sort took %s\n", elapsed)
	return s
}
