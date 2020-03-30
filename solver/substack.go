package main

type SubStacks struct {
	A []int
	B []int
}

func New(initialLen int) *SubStacks {
	ss := &SubStacks{
		A: make([]int, 1, 8),
		B: make([]int, 0, 8),
	}
	ss.A[0] = initialLen
	return ss
}

func (s *SubStacks) newSubStackA(i int) {
	s.A = append(s.A, i)
}

func (s *SubStacks) newSubStackB(i int) {
	s.B = append(s.B, i)
}

// CAR SubStacksA
func (s *SubStacks) lenFirstA() int {
	return s.A[len(s.A)-1]
}

// CDR SubStacksA
func (s *SubStacks) lenRestA() int {
	return sum(s.A[:len(s.A)-1])
}

func (s *SubStacks) decFirstA() {
	s.A[len(s.A)-1]--
}

// CAR SubStacksB
func (s *SubStacks) lenFirstB() int {
	return s.B[len(s.B)-1]
}

// CDR SubStacksB
func (s *SubStacks) lenRestB() int {
	return sum(s.B[:len(s.B)-1])
}

func (s *SubStacks) decFirstB() {
	s.B[len(s.B)-1]--
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}
