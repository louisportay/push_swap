package main

import (
	"fmt"
	"gitlab.com/louisportay/push_swap/stacks"
)

func lowest(indexer func(int) int, limit int) int {
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

func genOps(op string, n int, st stacks.Sorter) {
	ops := make([]string, n)
	for i := 0; i < n; i++ {
		ops[i] = op
	}
	st.AddOps(ops)
}

/*
	last element is chosen as pivot first,
	if no inferior element is found
	the previous is picked
*/

func Pivot(indexer func(int) int, subStackLen int, st stacks.Sorter) int {
	subStackLen--
	if indexer(subStackLen) == indexer(0) {
		return indexer(subStackLen)
	}
	for j := 0; j < subStackLen; j++ {
		if indexer(subStackLen) > indexer(j) {
			return indexer(subStackLen)
		}
	}
	return Pivot(indexer, subStackLen, st)
}

func n(f func(stacks.Sorter), n int, st stacks.Sorter) {
	for ; n > 0; n-- {
		f(st)
	}
}

func log(st stacks.Sorter, ss *SubStacks) {
	if ss != nil {
		fmt.Printf("SS.A: %v\nSS.B: %v\n", ss.A, ss.B)
	}
	st.PrintA()
	st.PrintB()
	st.PrintSorted()
	st.PrintOps()
}
