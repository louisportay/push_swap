package main

import (
	"gitlab.com/louisportay/push_swap/stacks"
)

func splitStackA(st stacks.Sorter, ss *SubStacks) int {
	p := Pivot(st.A, ss.lenFirstA())
	newSubStackLen := 0
	for i := 0; i < ss.lenFirstA()+newSubStackLen; i++ {
		if st.A(0) < p {
			pushB(st)
			newSubStackLen++
			ss.decFirstA()
		} else if st.A(0) >= p {
			rotateA(st)
		}
	}
	return newSubStackLen
}

// restores values which were previously rotated to the bottom of A
func restoreStackA(st stacks.Sorter, ss *SubStacks) {
	if st.LenSorted() > 0 || ss.lenRestA() > 0 {
		if ss.lenFirstA() > st.LenSorted()+ss.lenRestA() {
			w(rotateA, ss.lenRestA(), st)
			genOps("ra", st.LenSorted(), st)
		} else {
			w(revRotateA, ss.lenFirstA(), st)
		}

	}
}

func sortA(st stacks.Sorter, ss *SubStacks) {

	if ss.lenFirstA() <= 4 {
		sortTopA(ss.lenFirstA(), st)
		ss.A = ss.A[:len(ss.A)-1]
		return
	}

	ss.newSubStackB(splitStackA(st, ss))
	restoreStackA(st, ss)

	sortB(st, ss)
	sortA(st, ss)
}
