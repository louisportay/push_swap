package main

import "gitlab.com/louisportay/push_swap/stacks"

/*func swapBottomA(st stacks.Sorter) {
	st.RevRotateA()
	st.RevRotateA()
	st.SwapA()
	st.RotateA()
	st.RotateA()
	st.AddOps([]string{"rra", "rra", "sa", "ra", "ra"})
}
*/

func pushABSorted(st stacks.Sorter) {
	if st.A(0) < st.B(0) {
		pushASorted(st)
		pushBSorted(st)
	} else {
		pushBSorted(st)
		pushASorted(st)
	}
}

func pushASorted(st stacks.Sorter) {
	st.PushASorted()
	st.AddOps([]string{"ra"})
}

func pushBSorted(st stacks.Sorter) {
	st.PushBSorted()
	st.AddOps([]string{"pa", "ra"})
}

func pushA(st stacks.Sorter) {
	st.PushA()
	st.AddOps([]string{"pa"})
}

func pushB(st stacks.Sorter) {
	st.PushB()
	st.AddOps([]string{"pb"})
}

func swapA(st stacks.Sorter) {
	st.SwapA()
	st.AddOps([]string{"sa"})
}

func swapB(st stacks.Sorter) {
	st.SwapB()
	st.AddOps([]string{"sb"})
}

func swapBoth(st stacks.Sorter) {
	st.SwapBoth()
	st.AddOps([]string{"ss"})
}

func rotateA(st stacks.Sorter) {
	st.RotateA()
	st.AddOps([]string{"ra"})
}

func rotateB(st stacks.Sorter) {
	st.RotateB()
	st.AddOps([]string{"rb"})
}

func rotateBoth(st stacks.Sorter) {
	st.RotateBoth()
	st.AddOps([]string{"rr"})
}

func revRotateA(st stacks.Sorter) {
	st.RevRotateA()
	st.AddOps([]string{"rra"})
}

func revRotateB(st stacks.Sorter) {
	st.RevRotateB()
	st.AddOps([]string{"rrb"})
}

func revRotateBoth(st stacks.Sorter) {
	st.RevRotateBoth()
	st.AddOps([]string{"rrr"})
}
