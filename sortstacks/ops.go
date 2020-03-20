package sortstacks

func swap(i []int) {
	if len(i) > 1 {
		i[0], i[1] = i[1], i[0]
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
		var x int
		x, *a = (*a)[0], (*a)[1:]
		*b = append([]int{x}, *b...)
	}
}

func (s *sortStacks) PushA() {
	push(&s.b, &s.a)
}

func (s *sortStacks) PushB() {
	push(&s.a, &s.b)
}

func rotate(i *[]int) {
	if len(*i) > 1 {
		var x int
		x, *i = (*i)[0], (*i)[1:]
		*i = append(*i, x)
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

func revRotate(i *[]int) {
	if len(*i) > 1 {
		var x int
		x, *i = (*i)[len(*i)-1], (*i)[:len(*i)-1]
		*i = append([]int{x}, *i...)
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
