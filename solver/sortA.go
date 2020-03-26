package main

import (
	s "gitlab.com/louisportay/push_swap/sortstacks"
)

func splitStackA(st s.SortStacks, ss *SubStacks) (int, int) {
	p := Pivot(st.A, ss.lenFirstA(), st)
	newSubStackLen, alreadySorted := 0, 0
	for i := 0; i < ss.lenFirstA() + newSubStackLen; i++ {
		if st.FirstA() < p {
			pushB(st)
			newSubStackLen++
			ss.decFirstA()
		} else if st.FirstA() == p {
			pushB(st)
			newSubStackLen++
			ss.decFirstA()
			alreadySorted = ss.lenFirstA() + newSubStackLen - i - 1
		} else if st.FirstA() > p {
			rotateA(st)
		}
	}
	return newSubStackLen, alreadySorted
}

// restores values which were previously rotated to the bottom of A
func restoreStackA(st s.SortStacks, ss *SubStacks) {
	if st.LenSorted() > 0 || ss.lenRestA() > 0 {
		if ss.lenFirstA() > st.LenSorted() + ss.lenRestA() {
			n(rotateA, ss.lenRestA(), st)
			genOps("ra", st.LenSorted(), st)
		} else {
			n(revRotateA, ss.lenFirstA(), st)
		}
	}
}

func sortA(st s.SortStacks, ss *SubStacks) {

/*	if ss.lenFirstA() <= 4 {
		WIP
	}*/

	newSubStackLen, alreadySorted := splitStackA(st, ss)
	restoreStackA(st, ss)

	for ; alreadySorted > 0; alreadySorted-- {
		pushSorted(st)
		newSubStackLen--
	}

	ss.newSubStackB(newSubStackLen)
	sortB(st, ss)
	sortA(st, ss)
}
