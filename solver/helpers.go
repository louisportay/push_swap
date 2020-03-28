package main

import (
	"fmt"
	s "gitlab.com/louisportay/push_swap/sortstacks"
)

func lowest(indexer func (int) int, limit int) int {
	low_val, low_index := indexer(0), 0
	for i := 1; i < limit; i++ {
		if indexer(i) < low_val {
			low_val, low_index = indexer(i), i
		}
	}
	return low_index
}

func indexLowestValIntSlice(s []int) int {
	v, i := s[0], 0
	for j := 1; j < len(s); j++ {
		if s[j] < v {
			v, i = s[j], j
		}
	}
	return i
}

func genOps(op string, n int, st s.SortStacks) {
	ops := make([]string, n)
	for i:=0; i < n; i++ {
		ops[i] = op
	}
	st.AddOps(ops)
}

/*
	last element is chosen as pivot first,
	if no inferior element is found
	the previous is picked
*/

func Pivot(indexer func(int) int, subStackLen int, st s.SortStacks) int {
		subStackLen--
		if (indexer(subStackLen)  == indexer(0)) {
			return indexer(subStackLen)
		}
		for j := 0; j < subStackLen; j++ {
			if indexer(subStackLen) > indexer(j) {
				return indexer(subStackLen)
			}
		}
		return Pivot(indexer, subStackLen, st)
}

func n(f func(s.SortStacks), n int, st s.SortStacks) {
	for ; n > 0; n-- {
		f(st)
	}
}

func log(st s.SortStacks, ss *SubStacks) {
	if ss != nil {
		fmt.Printf("SS.A: %v\nSS.B: %v\n", ss.A, ss.B)
	}
	st.PrintA()
	st.PrintB()
	st.PrintSorted()
	st.PrintOps()
}
