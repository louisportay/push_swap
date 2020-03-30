package main

import (
	"gitlab.com/louisportay/push_swap/stacks"
)

func splitStackB(st stacks.Sorter, ss *SubStacks) (int, int) {
	p, lowstVal := Pivot(st.B, ss.lenFirstB())
	noRotate := consecutiveValuesHigherRef(st.B, ss.lenFirstB(), p)
	newSubStackLen, lowstValPast := 0, 0
	for i := 0; i < ss.lenFirstB()+newSubStackLen-noRotate+lowstValPast; i++ {
		if st.B(0) == lowstVal {
			pushBSorted(st)
			ss.decFirstB()
			lowstValPast = 1
		} else if st.B(0) < p {
			pushA(st)
			newSubStackLen++
			ss.decFirstB()
		} else if st.B(0) >= p {
			rotateB(st)
		}
	}
	return newSubStackLen, noRotate
}

func restoreStackB(st stacks.Sorter, ss *SubStacks, noRotate int) {
	if ss.lenRestB() > 0 {
		if ss.lenFirstB()-noRotate > ss.lenRestB()+noRotate {
			w(rotateB, ss.lenRestB()+noRotate, st)
		} else {
			w(revRotateB, ss.lenFirstB()-noRotate, st)
		}
	}
}

func sortB(st stacks.Sorter, ss *SubStacks) {

	if ss.lenFirstB() <= 4 {
		sortTopB(ss.lenFirstB(), st)
		ss.B = ss.B[:len(ss.B)-1]
		return
	}
	newSubStackLen, noRotate := splitStackB(st, ss)
	ss.newSubStackA(newSubStackLen)
	restoreStackB(st, ss, noRotate)

	sortA(st, ss)
	sortB(st, ss)
}
