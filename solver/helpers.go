package main

import (
	"fmt"//
	"os"//
	s "gitlab.com/louisportay/push_swap/sortstacks"
)

func genOps(op string, n int, st s.SortStacks) {
	ops := make([]string, n)
	for ; n > 0; n-- {
		ops[n] = op
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
	fmt.Printf("SS.A: %v\nSS.B: %v\n", ss.A, ss.B)
	st.PrintBoth()
	st.PrintSorted()
	st.PrintOps()
	os.Exit(0)
}
