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

func lowestRef(indexer func(int) int, limit int, ref int) (int, error) {
	i := 0
	for ; i < limit && indexer(i) <= indexer(ref); i++ {
	}
	if i == limit {
		return -1, fmt.Errorf("%v: highest value", indexer(ref))
	}
	v := indexer(i)
	for j := i + 1; j < limit; j++ {
		if indexer(j) > indexer(ref) && indexer(j) < v {
			v, i = indexer(j), j
		}
	}
	return i, nil
}

// return i such as s[i] > s[ref] && s[i] < s[*^ref]
func indexLowestValAfterRef(s []int, ref int) (int, error) {
	i := 0
	for ; i < len(s) && s[i] <= s[ref]; i++ {
	}
	if i == len(s) {
		return -1, fmt.Errorf("%v: highest value", s[ref])
	}
	v := s[i]
	for j := i + 1; j < len(s); j++ {
		if s[j] > s[ref] && s[j] < v {
			v, i = s[j], j
		}
	}
	return i, nil
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

/*
	pivot is median value
*/

func Pivot(indexer func(int) int, subStackLen int) int {
	p := lowest(indexer, subStackLen)
	for i := 0; i < subStackLen/2; i++ {
		p, _ = lowestRef(indexer, subStackLen, p)
	}
	return indexer(p)
}

// use with raw ops, methods
func m(f func(), n int, st stacks.Sorter) {
	for ; n > 0; n-- {
		f()
	}
}

// use with wrapped ops
func w(f func(stacks.Sorter), n int, st stacks.Sorter) {
	for ; n > 0; n-- {
		f(st)
	}
}

func genOps(op string, n int, st stacks.Sorter) {
	ops := make([]string, n)
	for i := 0; i < n; i++ {
		ops[i] = op
	}
	st.AddOps(ops)
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
