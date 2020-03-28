package main

import (
	"gitlab.com/louisportay/push_swap/stacks"
)

func sortTopB(subStackLen int, st stacks.Sorter) {
	simpleCases := func(lowestIndex int) {
		if lowestIndex == 0 {
			pushBSorted(st)
			sortTopB(subStackLen-1, st)
		} else if lowestIndex == 1 {
			swapB(st)
			pushBSorted(st)
			sortTopB(subStackLen-1, st)
		}
	}
	switch subStackLen {
	case 1:
		pushBSorted(st)
	case 2:
		if st.B(0) < st.B(1) {
			pushBSorted(st)
			pushBSorted(st)
		} else {
			swapB(st)
			pushBSorted(st)
			pushBSorted(st)
		}
	case 3:
		switch i := lowest(st.B, 3); i {
		case 0, 1:
			simpleCases(i)
		case 2:
			pushA(st)
			swapB(st)
			pushBSorted(st)
			pushABSorted(st)
		}
	case 4:
		switch i := lowest(st.B, 4); i {
		case 0, 1:
			simpleCases(i)
		case 2:
			pushA(st)
			swapB(st)
			pushBSorted(st)
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
		case 3:
			pushA(st)
			pushA(st)
			swapB(st)
			pushBSorted(st)
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
		}
	}
}
