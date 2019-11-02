package custom2

import (
	"fmt"
	"time"

	"github.com/ItsLeeOwen/gorting/pkg/util"
)

// Sort uses subdivide sampling
// and binary search to find insertion index
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
			s2 = util.Subdivide(s2, 0, len(s2)-1, v)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("custom2.Sort %s\n", elapsed)
	//fmt.Println(len(s2), s2)
	return s2
}
