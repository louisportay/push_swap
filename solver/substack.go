package main

// returns length of the current Sub Stack

func (s *Solver) lenSubStack() int {
	return s.lenSubStacks[len(s.lenSubStacks)-1]
}

// decrement length of the current substack

func (s *Solver) decLenSubStack() {
	s.lenSubStacks[len(s.lenSubStacks)-1]--
}

