package main

import (
	s "gitlab.com/louisportay/push_swap/sortstacks"
)

// OPTI: if pivot is the first element of the list, just push and rotate the whole
// substack, do not add any substacklen, and return

func splitStackB(st s.SortStacks, ss *SubStacks) (int, int) {
	p := Pivot(st.B, ss.lenFirstB(), st)
	newSubStackLen, alreadySorted := 0, 0
	for i := 0; i < ss.lenFirstB() + newSubStackLen; i++ {
		if st.FirstB() < p {
			pushA(st)
			newSubStackLen++
			ss.decFirstB()
		} else if st.FirstB() == p {
			pushA(st)
			newSubStackLen++
			ss.decFirstB()
			alreadySorted = ss.lenFirstB() + newSubStackLen - i - 1
		} else if st.FirstB() > p {
			rotateB(st)
		}
	}
	return newSubStackLen, alreadySorted
}

func restoreStackB(st s.SortStacks, ss *SubStacks) {
	if ss.lenRestB() > 0 {
		if ss.lenFirstB() > ss.lenRestB() {
			n(rotateB, ss.lenRestB(), st)
		} else {
			n(revRotateB, ss.lenFirstB(), st)
		}
	}
}

func sortB(st s.SortStacks, ss *SubStacks) {

	if ss.lenFirstB() <= 4 {
		sortTopB(ss.lenFirstB(), st)
		ss.B = ss.B[:len(ss.B)-1]
		return
	}

	newSubStackLen, alreadySorted := splitStackB(st, ss)
	restoreStackB(st, ss)

	for ; alreadySorted > 0; alreadySorted-- {
		rotateSorted(st)
		newSubStackLen--
	}

	ss.newSubStackA(newSubStackLen)

	sortA(st, ss)
	sortB(st, ss)
}
