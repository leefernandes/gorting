package util

import (
	"math/rand"
	"time"
)

func CopySlice(s1 []int) []int {
	s2 := make([]int, len(s1))
	copy(s2, s1)
	return s2
}

func GenerateRandomSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(10000001) // - rand.Intn(10000001)
	}
	return slice
}

// FindInsertion does a binary search for insertion index of v
func FindInsertion(s []int, start int, end int, v int) int {
	l := end - start
	if 0 == l {
		if v < s[start] {
			return start
		}
		return start + 1
	}
	mI := start + (end-start)/2
	m := s[mI]

	if v == m {
		return mI
	} else if v < m {
		return FindInsertion(s, start, mI, v)
	}
	return FindInsertion(s, mI+1, end, v)
}

// FindMinMax returns the min/max values of a slice
func FindMinMax(s []int, start, end int) (min, max int) {
	min = s[start]
	max = min
	for i := start + 1; i <= end; i++ {
		v := s[i]
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return
}

func InsertAt(s []int, i int, v int) []int {
	//s = append(s[:i], append([]int{v}, s[i:]...)...)
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = v
	return s
}

// SpotAhead samples a sorted slice
// in 8ths until finding a segment
// where the value should be inserted
// and returns the start/end of the
// sample size
func SpotAhead(s []int, start int, end int, v int) (int, int) {
	d := (end - start) / 8
	// sampling
	for i := start + d; i < end; i += d {
		if s[i] >= v {
			return i - d, i
		}
	}

	for i := end; i > start; i-- {
		c := s[i]
		if v >= c {
			return i, end
		}
	}
	return start, start
}

// Subdivide uses spotAhead sampling
// on a sorted slice until it has
// size < 16 then does a binary search
func Subdivide(s []int, start int, end int, v int) []int {
	if end-start < 16 {
		i := FindInsertion(s, start, end, v)
		s = InsertAt(s, i, v)

		// for i := start; i <= end; i++ {
		// 	c := s[i]
		// 	if v <= c {
		// 		s = insertAt(s, i, v)
		// 		break
		// 	}
		// 	if i == end {
		// 		s = insertAt(s, 0, v)
		// 	}
		// }

	} else {
		s2, e2 := SpotAhead(s, start, end, v)
		s = Subdivide(s, s2, e2, v)
	}
	return s
}
