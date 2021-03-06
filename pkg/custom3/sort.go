package custom3

import (
	"fmt"
	"sync"
	"time"

	"github.com/leefernandes/gorting/pkg/quick"
	"github.com/leefernandes/gorting/pkg/util"
)

// Sort partitions a slice into 4 blocks
// and runs a quicksort goroutine for each
func Sort(s []int) []int {
	start := time.Now()
	l := len(s)
	p := 4
	d := l / p
	pEnd := p - 1
	min := s[0]
	max := min
	var rng, mid int

	var wg sync.WaitGroup
	wg.Add(p)

	for i := 0; i < p; i++ {
		start := i * d
		var end int
		if i == pEnd {
			end = l - 1
		} else {
			end = start + d
		}
		go func() {
			lMin, lMax := util.FindMinMax(s, start, end)
			if lMin < min {
				min = lMin
			}
			if lMax > max {
				max = lMax
			}
			wg.Done()
		}()
	}

	wg.Wait()

	rng = max - min
	mid = min + (rng)/2

	q1, q2, q3, q4 := quadrantize(s, min, max, rng, mid)

	starts := []int{0, q1, q2, q3}
	ends := []int{q1, q2, q3, q4}

	wg.Add(p)

	for i := 0; i < p; i++ {
		start := starts[i]
		end := ends[i]
		go func() {
			quick.Quicksort(s, start, end)
			wg.Done()
		}()
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("custom3.Sort %s\n", elapsed)
	//	fmt.Println(len(s), s)
	return s
}

// quadrantize partitions a slice into 4 blocks
// and returns the end pivot of each blocks
func quadrantize(s []int, min, max, rng, mid int) (int, int, int, int) {
	l := len(s)
	q := rng / 4

	q1 := min + q
	q2 := q1 + q
	q3 := q2 + q

	var qp1, qp2, qp3, qp4 int

	for i := 0; i < l; i++ {
		v := s[i]
		if v < q1 {
			s[i], s[qp3], s[qp2], s[qp1] = s[qp3], s[qp2], s[qp1], v
			qp1++
			qp2++
			qp3++
			qp4++
		} else if v < q2 {
			s[i], s[qp3], s[qp2] = s[qp3], s[qp2], v
			qp2++
			qp3++
			qp4++
		} else if v < q3 {
			s[i], s[qp3] = s[qp3], v
			qp3++
			qp4++
		} else {
			s[i], s[qp4] = s[qp4], v
			qp4++
		}
	}

	return qp1, qp2, qp3, qp4
}
