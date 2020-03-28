package main

import (
	s "gitlab.com/louisportay/push_swap/sortstacks"
)

func sortTopB(subStackLen int, st s.SortStacks) {
	simpleCases := func(lowestIndex int) {
		if lowestIndex == 0 {
			pushSorted(st)
			sortTopB(subStackLen-1, st)
		} else if lowestIndex == 1 {
			swapB(st); pushSorted(st)
			sortTopB(subStackLen-1, st)
		}
	}
	switch subStackLen {
	case 1:
		pushSorted(st)
	case 2:
		if st.B(0) < st.B(1) {
			pushSorted(st); pushSorted(st)
		} else {
			swapB(st); pushSorted(st); pushSorted(st)
		}
	case 3:
		switch i := lowest(st.B, 3); i {
		case 0, 1:
			simpleCases(i)
		case 2:
			pushA(st); swapB(st); pushSorted(st); pushAndRotateSorted(st)
		}
	case 4:
		switch i := lowest(st.B, 4); i {
		case 0, 1:
			simpleCases(i)
		case 2:
			pushA(st); swapB(st); pushSorted(st)
			switch indexLowestValIntSlice([]int{st.A(0), st.B(0), st.B(1)}) {
			case 0:
				rotateSorted(st); sortTopB(2, st)
			case 1:
				pushSorted(st); pushAndRotateSorted(st)
			case 2:
				swapB(st); pushSorted(st); pushAndRotateSorted(st)
			}
		case 3:
			pushA(st); pushA(st); swapB(st); pushSorted(st)
			switch indexLowestValIntSlice([]int{st.B(0), st.A(0), st.A(1)}) {
			case 0:
				pushSorted(st); sortTopA(2, st)
			case 1:
				rotateSorted(st); pushAndRotateSorted(st)
			case 2:
				swapA(st); rotateSorted(st); pushAndRotateSorted(st)
			}
		}
	}
}
