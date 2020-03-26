package main

import (
	s "gitlab.com/louisportay/push_swap/sortstacks"
)

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

	// if lenFirstB <= 4 --> special case, triggers end of recursion

	newSubStackLen, alreadySorted := splitStackB(st, ss)
	restoreStackB(st, ss)

	for ; alreadySorted > 0; alreadySorted-- {
		rotateSorted(st)
		newSubStackLen--
	}

	ss.newSubStackA(newSubStackLen)

	log(st, ss)
	sortA(st, ss)
	sortB(st, ss)
}
