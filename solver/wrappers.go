package main

import s "gitlab.com/louisportay/push_swap/sortstacks"

/*func swapBottomA(st s.SortStacks) {
	st.RevRotateA()
	st.RevRotateA()
	st.SwapA()
	st.RotateA()
	st.RotateA()
	st.AddOps([]string{"rra", "rra", "sa", "ra", "ra"})
}
*/

func pushAndRotateSorted(st s.SortStacks) {
	if st.FirstA() < st.FirstB() {
		rotateSorted(st); pushSorted(st)
	} else {
		pushSorted(st); rotateSorted(st)
	}
}

func rotateSorted(st s.SortStacks) {
	st.RotateSorted()
	st.AddOps([]string{"ra"})
}

func pushSorted(st s.SortStacks) {
	st.PushSorted()
	st.AddOps([]string{"pa", "ra"})
}

func pushA(st s.SortStacks) {
	st.PushA()
	st.AddOps([]string{"pa"})
}

func pushB(st s.SortStacks) {
	st.PushB()
	st.AddOps([]string{"pb"})
}

func swapA(st s.SortStacks) {
	st.SwapA()
	st.AddOps([]string{"sa"})
}

func swapB(st s.SortStacks) {
	st.SwapB()
	st.AddOps([]string{"sb"})
}

func swapBoth(st s.SortStacks) {
	st.SwapBoth()
	st.AddOps([]string{"ss"})
}

func rotateA(st s.SortStacks) {
	st.RotateA()
	st.AddOps([]string{"ra"})
}

func rotateB(st s.SortStacks) {
	st.RotateB()
	st.AddOps([]string{"rb"})
}

func rotateBoth(st s.SortStacks) {
	st.RotateBoth()
	st.AddOps([]string{"rr"})
}

func revRotateA(st s.SortStacks) {
	st.RevRotateA()
	st.AddOps([]string{"rra"})
}

func revRotateB(st s.SortStacks) {
	st.RevRotateB()
	st.AddOps([]string{"rrb"})
}

func revRotateBoth(st s.SortStacks) {
	st.RevRotateBoth()
	st.AddOps([]string{"rrr"})
}
