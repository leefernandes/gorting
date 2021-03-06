package custom1

import (
	"fmt"
	"time"

	"github.com/leefernandes/gorting/pkg/util"
)

// Sort runs a slightly optimized O(n2)
// moving values into a sorted slice starting
// comparison at either the left or right
func Sort(s []int) []int {
	start := time.Now()

	l := len(s)
	min := s[0]
	max := min
	s2 := make([]int, 0, l)
	s2 = append(s2, min)
	for i := 1; i < l; i++ {
		v := s[i]
		if v <= min {
			min = v
			//s2 = append([]int{v}, s2...)
			s2 = util.InsertAt(s2, 0, v)

		} else if v >= max {
			max = v
			s2 = append(s2, v)
		} else {
			l2 := len(s2)
			m := s2[l2/2]
			l2End := l2 - 1

			if v <= m {
				// iterate from left
				for i2 := 0; i2 < l2; i2++ {
					v2 := s2[i2]
					if v <= v2 {
						s2 = util.InsertAt(s2, i2, v)
						break
					}
					if i2 == l2End {
						//s2 = append([]int{v}, s2...)
						s2 = util.InsertAt(s2, 0, v)
					}
				}
			} else {
				// iterate from right
				for i2 := l2End; i2 >= 0; i2-- {
					v2 := s2[i2]
					if v >= v2 {
						s2 = util.InsertAt(s2, i2+1, v)
						break
					}
					if i2 == 0 {
						s2 = append(s2, v)
					}
				}
			}

		}
	}

	elapsed := time.Since(start)
	fmt.Printf("custom1.Sort %s\n", elapsed)
	//fmt.Println(len(s2), s2)
	return s2
}
