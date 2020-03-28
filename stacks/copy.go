package stacks

/*
	B list must be empty and
	A is first reversed and then concatenated with
	the 'sorted' special list
*/

func (st *sortStacks) virtualA() []int {
	cp := st.CopyA()
	cp = append(cp, st.CopySorted()...)
	return cp
}

func (st *sortStacks) CopyA() []int {
	cp := make([]int, len(st.a))
	copy(cp, st.a)
	reverseInts(cp)
	return cp
}

func (st *sortStacks) CopyB() []int {
	cp := make([]int, len(st.b))
	copy(cp, st.b)
	reverseInts(cp)
	return cp
}

func (st *sortStacks) CopySorted() []int {
	cp := make([]int, len(st.sorted))
	copy(cp, st.sorted)
	return cp
}
