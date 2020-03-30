package main

import (
	"fmt"
	"os"
	"gitlab.com/louisportay/push_swap/stacks"
)

func splitStackA(st stacks.Sorter, ss *SubStacks) (int, int, Order) {
	p, order := Pivot(st.A, ss.lenFirstA(), None, st)
	newSubStackLen, belowPivot:= 0, 0
	for i := 0; i < ss.lenFirstA()+newSubStackLen; i++ {
		if st.A(0) < p {
			pushB(st)
			newSubStackLen++
			ss.decFirstA()
		} else if st.A(0) == p {
		//	pushB(st)
		//	newSubStackLen++
		//	ss.decFirstA()
			belowPivot = ss.lenFirstA() + newSubStackLen - i// - 1
			if order == Descending {
				break ;
			}
		} else if st.A(0) > p {
			rotateA(st)
		}
	}
	return newSubStackLen, belowPivot, order
}

// restores values which were previously rotated to the bottom of A
func restoreStackA(st stacks.Sorter, ss *SubStacks, belowPivot int) {
	if st.LenSorted() > 0 || ss.lenRestA() > 0 {
		if ss.lenFirstA() > st.LenSorted()+ss.lenRestA()+belowPivot {
			w(rotateA, ss.lenRestA()+belowPivot, st)
			genOps("ra", st.LenSorted(), st)
		} else {
			w(revRotateA, ss.lenFirstA()-belowPivot, st)
		}

	}
}

func sortA(st stacks.Sorter, ss *SubStacks) {

	if ss.lenFirstA() <= 4 {
		sortTopA(ss.lenFirstA(), st)
		ss.A = ss.A[:len(ss.A)-1]
		return
	}

	newSubStackLen, belowPivot, order := splitStackA(st, ss)

	if order == Descending {
		fmt.Fprintf(os.Stderr, "A Asc: %v\n", belowPivot)
		restoreStackA(st, ss, belowPivot)
	} else {
		restoreStackA(st, ss, 0)
		fmt.Fprintf(os.Stderr, "A Asc: %v\n", belowPivot)
		for ; belowPivot > 0; belowPivot-- {
			pushBSorted(st)
			newSubStackLen--
		}
	}

	ss.newSubStackB(newSubStackLen)
	sortB(st, ss)
	sortA(st, ss)
}
