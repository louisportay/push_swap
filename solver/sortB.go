package main

import (
	"gitlab.com/louisportay/push_swap/stacks"
)

func splitStackB(st stacks.Sorter, ss *SubStacks) int {
	p := Pivot(st.B, ss.lenFirstB())
	newSubStackLen := 0
	for i := 0; i < ss.lenFirstB()+newSubStackLen; i++ {
		if st.B(0) < p {
			pushA(st)
			newSubStackLen++
			ss.decFirstB()
		} else if st.B(0) >= p {
			rotateB(st)
		}
	}
	return newSubStackLen
}

func restoreStackB(st stacks.Sorter, ss *SubStacks) {
	if ss.lenRestB() > 0 {
		if ss.lenFirstB() > ss.lenRestB() {
			w(rotateB, ss.lenRestB(), st)
		} else {
			w(revRotateB, ss.lenFirstB(), st)
		}
	}
}

func sortB(st stacks.Sorter, ss *SubStacks) {

	if ss.lenFirstB() <= 4 {
		sortTopB(ss.lenFirstB(), st)
		ss.B = ss.B[:len(ss.B)-1]
		return
	}

	ss.newSubStackA(splitStackB(st, ss))
	restoreStackB(st, ss)

	sortA(st, ss)
	sortB(st, ss)
}
