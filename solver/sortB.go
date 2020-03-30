package main

import (
	"gitlab.com/louisportay/push_swap/stacks"
	"os"
	"fmt"
)

func splitStackB(st stacks.Sorter, ss *SubStacks) (int, int, Order) {
	p, order := Pivot(st.B, ss.lenFirstB(), None, st)
	newSubStackLen, belowPivot := 0, 0
	for i := 0; i < ss.lenFirstB()+newSubStackLen; i++ {
		if st.B(0) < p {
			pushA(st)
			newSubStackLen++
			ss.decFirstB()
		} else if st.B(0) == p {
			pushA(st)
			newSubStackLen++
			ss.decFirstB()
			belowPivot = ss.lenFirstB() + newSubStackLen - i - 1
			if order == Descending {
				break ;
			}
		} else if st.B(0) > p {
			rotateB(st)
		}
	}
	return newSubStackLen, belowPivot, order
}

func restoreStackB(st stacks.Sorter, ss *SubStacks, belowPivot int) {
	if ss.lenRestB() > 0 {
		if ss.lenFirstB() > ss.lenRestB()+belowPivot {
			w(rotateB, ss.lenRestB()+belowPivot, st)
		} else {
			w(revRotateB, ss.lenFirstB()-belowPivot, st)
		}
	}
}

func sortB(st stacks.Sorter, ss *SubStacks) {

	if ss.lenFirstB() <= 4 {
		sortTopB(ss.lenFirstB(), st)
		ss.B = ss.B[:len(ss.B)-1]
		return
	}

	newSubStackLen, belowPivot, order := splitStackB(st, ss)
	if order == Descending {
		fmt.Fprintf(os.Stderr, "B Desc: %v\n", belowPivot)
		restoreStackA(st, ss, belowPivot)
	} else {
		restoreStackA(st, ss, 0)
		fmt.Fprintf(os.Stderr, "B Asc: %v\n", belowPivot)
		for ; belowPivot > 0; belowPivot-- {
			pushASorted(st)
			newSubStackLen--
		}
	}

	ss.newSubStackA(newSubStackLen)

	sortA(st, ss)
	sortB(st, ss)
}
