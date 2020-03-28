package sortstacks

func swap(i []int) {
	if len(i) > 1 {
		i[len(i)-1], i[len(i)-1-1] = i[len(i)-1-1], i[len(i)-1]
	}
}

func (s *sortStacks) SwapA() {
	swap(s.a)
}

func (s *sortStacks) SwapB() {
	swap(s.b)
}

func (s *sortStacks) SwapBoth() {
	swap(s.a)
	swap(s.b)
}

// first element of a go on top of b

func push(a *[]int, b *[]int) {
	if len(*a) > 0 {
		x := (*a)[len(*a)-1]
		*a = (*a)[:len(*a)-1]
		*b = append(*b, x)
	}
}

func (s *sortStacks) PushA() {
	push(&s.b, &s.a)
}

func (s *sortStacks) PushB() {
	push(&s.a, &s.b)
}

// Expensive op
func rotate(i *[]int) {
	if len(*i) > 1 {
		x := (*i)[len(*i)-1]
		*i = append([]int{x}, (*i)[:len(*i)-1]...)
	}
}

func (s *sortStacks) RotateA() {
	rotate(&s.a)
}

func (s *sortStacks) RotateB() {
	rotate(&s.b)
}

func (s *sortStacks) RotateBoth() {
	rotate(&s.a)
	rotate(&s.b)
}

func (s *sortStacks) pop(i *[]int) {
	x := (*i)[len(*i)-1]
	*i = (*i)[:len(*i)-1]
	s.sorted = append(s.sorted, x)
}

func (s *sortStacks) RotateSorted() {
	s.pop(&s.a)
}

func (s *sortStacks) PushSorted() {
	s.pop(&s.b)
}

func revRotate(i *[]int) {
	if len(*i) > 1 {
		x := (*i)[0]
		*i = append((*i)[1:], x)
	}
}

func (s *sortStacks) RevRotateA() {
	revRotate(&s.a)
}

func (s *sortStacks) RevRotateB() {
	revRotate(&s.b)
}

func (s *sortStacks) RevRotateBoth() {
	revRotate(&s.a)
	revRotate(&s.b)
}


// To Delete
func (st *sortStacks) FirstA() int {
	return st.a[len(st.a)-1]
}

// To Delete
func (st *sortStacks) FirstB() int {
	return st.b[len(st.b)-1]
}

// To Delete Maybe
func (st *sortStacks) LastA() int {
	return st.a[0]
}

// To Delete Maybe
func (st *sortStacks) LastB() int {
	return st.b[0]
}

