package main

import (
	"gitlab.com/louisportay/push_swap/parser"
	"gitlab.com/louisportay/push_swap/stacks"
	"os"
)

func upperBound(a []int, lowerBound int) int {
	up, _ := indexLowestValAfterRef(a, lowerBound)
	if up == lowerBound+1 {
		for n, _ := indexLowestValAfterRef(a, up); n == up+1; n, _ = indexLowestValAfterRef(a, up) {
			up = n
		}
	} else {
		up = lowerBound
	}
	return up
}

func initialSorted(st stacks.Sorter) {
	a := st.CopyA()
	low := indexLowestValIntSlice(a)
	up := upperBound(a, low)
	if low == len(a)-1 || up == low {
		return
	}
	if len(a[low:up+1]) == len(a) {
		st.MoveASorted(len(a))
	} else if low == 0 {
		m(st.PushASorted, up+1, st)
		if len(a)-up-1 >= up+1 {
			genOps("ra", up+1, st)
		} else {
			genOps("rra", len(a)-up-1, st)
		}
	} else if up+1 >= len(a)-up-1 && up-low+1 >= len(a)-up-1 {
		w(revRotateA, len(a)-up-1, st)
		st.MoveASorted(up-low+1)
	} else if up+1 < len(a)-up-1 && up-low+1 >= low {
		w(rotateA, low, st)
		w(pushASorted, up-low+1, st)
	}
}

func main() {
	st := stacks.NewSorter(parser.BuildStack(os.Args[1:]))
	if st.LenA() == 1 {
		os.Exit(0)
	}
	initialSorted(st)
	ss := New(st.LenA())
	sortA(st, ss)
	st.PrintOps()
}
