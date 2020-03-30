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

// return i such as s[i] > s[ref] && s[i] < s[*^ref]
func indexLowestValAfterRef(s []int, ref int) (int, error) {
	i := 0
	for ; i < len(s) && s[i] <= s[ref]; i++ { ; }
	if i == len(s) {
		return -1, fmt.Errorf("%v: highest value in %v", ref, s)
	}
	v := s[i]
	for j := i+1; j < len(s); j++ {
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

type Order int
const (
	None = iota
	Ascending
	Descending
)

func Pivot(indexer func(int) int, subStackLen int, o Order, st stacks.Sorter) (int, Order) {
	subStackLen--
	if indexer(subStackLen) == indexer(0) {
		return indexer(subStackLen), o
	}
	inf, sup := false, false
	for i := 0; i < subStackLen; i++ {
		if indexer(i) < indexer(subStackLen) {
			inf = true
		} else if indexer(i) > indexer(subStackLen) {
			sup = true
		}
		if sup == true && inf == true {
			break ;
		}
	}
	if inf == true && sup == false && (o == None || o == Descending) {
			return Pivot(indexer, subStackLen, Descending, st)
	} else if inf == false && sup == true && (o == None || o == Ascending) {
			return Pivot(indexer, subStackLen, Ascending, st)
	} else {
			return indexer(subStackLen), o
	}
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

func log(st stacks.Sorter, ss *SubStacks) {
	if ss != nil {
		fmt.Printf("SS.A: %v\nSS.B: %v\n", ss.A, ss.B)
	}
	st.PrintA()
	st.PrintB()
	st.PrintSorted()
	st.PrintOps()
}
