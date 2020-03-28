package main

import (
	"gitlab.com/louisportay/push_swap/stacks"
)

func splitStackA(st stacks.Sorter, ss *SubStacks) (int, int) {
	p := Pivot(st.A, ss.lenFirstA(), st)
	newSubStackLen, alreadySorted := 0, 0
	for i := 0; i < ss.lenFirstA()+newSubStackLen; i++ {
		if st.A(0) < p {
			pushB(st)
			newSubStackLen++
			ss.decFirstA()
		} else if st.A(0) == p {
			pushB(st)
			newSubStackLen++
			ss.decFirstA()
			alreadySorted = ss.lenFirstA() + newSubStackLen - i - 1
		} else if st.A(0) > p {
			rotateA(st)
		}
	}
	return newSubStackLen, alreadySorted
}

// restores values which were previously rotated to the bottom of A
func restoreStackA(st stacks.Sorter, ss *SubStacks) {
	if st.LenSorted() > 0 || ss.lenRestA() > 0 {
		if ss.lenFirstA() > st.LenSorted()+ss.lenRestA() {
			n(rotateA, ss.lenRestA(), st)
			genOps("ra", st.LenSorted(), st)
		} else {
			n(revRotateA, ss.lenFirstA(), st)
		}
	}
}

func sortA(st stacks.Sorter, ss *SubStacks) {

	if ss.lenFirstA() <= 4 {
		sortTopA(ss.lenFirstA(), st)
		ss.A = ss.A[:len(ss.A)-1]
		return
	}

	newSubStackLen, alreadySorted := splitStackA(st, ss)
	restoreStackA(st, ss)

	for ; alreadySorted > 0; alreadySorted-- {
		pushBSorted(st)
		newSubStackLen--
	}

	ss.newSubStackB(newSubStackLen)
	sortB(st, ss)
	sortA(st, ss)
}
