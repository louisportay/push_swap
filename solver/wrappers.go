package main

import ss "gitlab.com/louisportay/push_swap/sortstacks"

func swapBottomA(st ss.SortStacks) {
	st.RevRotateA()
	st.RevRotateA()
	st.SwapA()
	st.RotateA()
	st.RotateA()
	st.AddOps([]string{"rra", "rra", "sa", "ra", "ra"})
}

/*
func pushBottomA(st ss.SortStacks) {
	st.PushA()
	st.RotateA()
	st.AddOps([]string{"pa", "ra"})
}*/

func pushA(st ss.SortStacks) {
	st.PushA()
	st.AddOps([]string{"pa"})
}

func pushB(st ss.SortStacks) {
	st.PushB()
	st.AddOps([]string{"pb"})
}

func swapA(st ss.SortStacks) {
	st.SwapA()
	st.AddOps([]string{"sa"})
}

func swapB(st ss.SortStacks) {
	st.SwapB()
	st.AddOps([]string{"sb"})
}

func swapBoth(st ss.SortStacks) {
	st.SwapBoth()
	st.AddOps([]string{"ss"})
}

func rotateA(st ss.SortStacks) {
	st.RotateA()
	st.AddOps([]string{"ra"})
}

func rotateB(st ss.SortStacks) {
	st.RotateB()
	st.AddOps([]string{"rb"})
}

func rotateBoth(st ss.SortStacks) {
	st.RotateBoth()
	st.AddOps([]string{"rr"})
}

func revRotateA(st ss.SortStacks) {
	st.RevRotateA()
	st.AddOps([]string{"rra"})
}

func revRotateB(st ss.SortStacks) {
	st.RevRotateB()
	st.AddOps([]string{"rrb"})
}

func revRotateBoth(st ss.SortStacks) {
	st.RevRotateBoth()
	st.AddOps([]string{"rrr"})
}

func rotateSorted(st ss.SortStacks) {
	st.RotateSorted()
	st.AddOps([]string{"ra"})
}

func pushSorted(st ss.SortStacks) {
	st.PushSorted()
	st.AddOps([]string{"pa", "ra"})
}
