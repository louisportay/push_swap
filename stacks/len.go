package stacks

func (st *sortStacks) LenA() int {
	return len(st.a)
}

func (st *sortStacks) LenB() int {
	return len(st.b)
}

func (st *sortStacks) LenSorted() int {
	return len(st.sorted)
}
