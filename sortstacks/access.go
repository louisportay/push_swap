package sortstacks

func (st *sortStacks) FirstA() int {
	if len(st.a) == 0 {
		panic("Illegal Access to first Element of A")
	}
	return st.a[len(st.a)-1]
}

func (st *sortStacks) FirstB() int {
	if len(st.b) == 0 {
		panic("Illegal Access to first Element of B")
	}
	return st.b[len(st.b)-1]
}

func (st *sortStacks) LastA() int {
	if len(st.a) == 0 {
		panic("Illegal Access to last Element of A")
	}
	return st.a[0]
}

func (st *sortStacks) LastB() int {
	if len(st.b) == 0 {
		panic("Illegal Access to last Element of B")
	}
	return st.b[0]
}
