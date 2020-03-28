package main

import (
	"gitlab.com/louisportay/push_swap/stacks"
)

func sortTopA(subStackLen int, st stacks.Sorter) {
	simpleCases := func(lowestIndex int) {
		if lowestIndex == 0 {
			pushASorted(st)
			sortTopA(subStackLen-1, st)
		} else if lowestIndex == 1 {
			swapA(st)
			pushASorted(st)
			sortTopA(subStackLen-1, st)
		}
	}
	switch subStackLen {
	case 1:
		pushASorted(st)
	case 2:
		if st.A(0) < st.A(1) {
			pushASorted(st)
			pushASorted(st)
		} else {
			swapA(st)
			pushASorted(st)
			pushASorted(st)
		}
	case 3:
		switch i := lowest(st.A, 3); i {
		case 0, 1:
			simpleCases(i)
		case 2:
			pushB(st)
			swapA(st)
			pushASorted(st)
			pushABSorted(st)
		}
	case 4:
		switch i := lowest(st.A, 4); i {
		case 0, 1:
			simpleCases(i)
		case 2:
			pushB(st)
			swapA(st)
			pushASorted(st)
			switch indexLowestValIntSlice([]int{st.B(0), st.A(0), st.A(1)}) {
			case 0:
				pushBSorted(st)
				sortTopA(2, st)
			case 1:
				pushASorted(st)
				pushABSorted(st)
			case 2:
				swapA(st)
				pushASorted(st)
				pushABSorted(st)
			}
		case 3:
			pushB(st)
			pushB(st)
			swapA(st)
			pushASorted(st)
			switch indexLowestValIntSlice([]int{st.A(0), st.B(0), st.B(1)}) {
			case 0:
				pushASorted(st)
				sortTopB(2, st)
			case 1:
				pushBSorted(st)
				pushABSorted(st)
			case 2:
				swapB(st)
				pushBSorted(st)
				pushABSorted(st)
			}
		}
	}
}
