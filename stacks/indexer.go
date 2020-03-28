package stacks

func (st *sortStacks) A(i int) int {
	return st.a[len(st.a)-1-i]
}

func (st *sortStacks) B(i int) int {
	return st.b[len(st.b)-1-i]
}

func (st *sortStacks) Sorted(i int) int {
	return st.sorted[i]
}

func (st *sortStacks) Op(op string) Op {
	return st.f[op]
}

// To Delete Maybe
func (st *sortStacks) LastA() int {
	return st.a[0]
}

// To Delete Maybe
func (st *sortStacks) LastB() int {
	return st.b[0]
}
