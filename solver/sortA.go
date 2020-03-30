package main

import (
	"gitlab.com/louisportay/push_swap/stacks"
)

func splitStackA(st stacks.Sorter, ss *SubStacks) (int, int) {
	p, _ := Pivot(st.A, ss.lenFirstA())
	noRotate := consecutiveValuesHigherRef(st.A, ss.lenFirstA(), p)
	newSubStackLen := 0
	for i := 0; i < ss.lenFirstA()+newSubStackLen-noRotate; i++ {
/*		if st.A(0) == lowstVal {
			pushASorted(st)
			ss.decFirstA()
		} else*/ if st.A(0) < p {
			pushB(st)
			newSubStackLen++
			ss.decFirstA()
		} else if st.A(0) >= p {
			rotateA(st)
		}
	}
	return newSubStackLen, noRotate
}

// restores values which were previously rotated to the bottom of A
func restoreStackA(st stacks.Sorter, ss *SubStacks, noRotate int) {
	if st.LenSorted() > 0 || ss.lenRestA() > 0 {
//		fmt.Fprintf(os.Stderr, "%v\n", noRotate)
		if ss.lenFirstA()-noRotate > st.LenSorted()+ss.lenRestA()+noRotate {
			w(rotateA, ss.lenRestA()+noRotate, st)
			genOps("ra", st.LenSorted(), st)
		} else {
			w(revRotateA, ss.lenFirstA()-noRotate, st)
		}

	}
}

func sortA(st stacks.Sorter, ss *SubStacks) {

	if ss.lenFirstA() <= 4 {
		sortTopA(ss.lenFirstA(), st)
		ss.A = ss.A[:len(ss.A)-1]
		return
	}
	newSubStackLen, noRotate := splitStackA(st, ss)
	ss.newSubStackB(newSubStackLen)
	restoreStackA(st, ss, noRotate)

	sortB(st, ss)
	sortA(st, ss)
}
