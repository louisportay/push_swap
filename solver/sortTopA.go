package main

import (
	s "gitlab.com/louisportay/push_swap/sortstacks"
)

func sortTopA(subStackLen int, st s.SortStacks) {
	simpleCases := func(lowestIndex int) {
		if lowestIndex == 0 {
			rotateSorted(st)
			sortTopA(subStackLen-1, st)
		} else if lowestIndex == 1 {
			swapA(st); rotateSorted(st)
			sortTopA(subStackLen-1, st)
		}
	}
	switch subStackLen {
	case 1:
		rotateSorted(st)
	case 2:
		if st.A(0) < st.A(1) {
			rotateSorted(st); rotateSorted(st)
		} else {
			swapA(st); rotateSorted(st); rotateSorted(st)
		}
	case 3:
		switch i := lowest(st.A, 3); i {
		case 0, 1:
			simpleCases(i)
		case 2:
			pushB(st); swapA(st); rotateSorted(st); pushAndRotateSorted(st)
		}
	case 4:
		switch i := lowest(st.A, 4); i {
		case 0, 1:
			simpleCases(i)
		case 2:
			pushB(st); swapA(st); rotateSorted(st)
			switch indexLowestValIntSlice([]int{st.B(0), st.A(0), st.A(1)}) {
			case 0:
				pushSorted(st); sortTopA(2, st)
			case 1:
				rotateSorted(st); pushAndRotateSorted(st)
			case 2:
				swapA(st); rotateSorted(st); pushAndRotateSorted(st)
			}
		case 3:
			pushB(st); pushB(st); swapA(st); rotateSorted(st)
			switch indexLowestValIntSlice([]int{st.A(0), st.B(0), st.B(1)}) {
			case 0:
				rotateSorted(st); sortTopB(2, st)
			case 1:
				pushSorted(st); pushAndRotateSorted(st)
			case 2:
				swapB(st); pushSorted(st); pushAndRotateSorted(st)
			}
		}
	}
}
