package main

import (
	"os"
	"fmt"
	"gitlab.com/louisportay/push_swap/parser"
	ss "gitlab.com/louisportay/push_swap/sortstacks"
)

type Solver struct {
	lenSubStacks []int
	newSubStackLen int
	//RotateStack []int
	rotations int
}

func pushSubStackA(st ss.SortStacks, so *Solver) {
	pushA(st)
	so.newSubStackLen++
	so.decLenSubStack()
}

// 'left' and 'right' are indexes, not the underlying values


// OPTI: si tous les elements de la pile sont inferieurs a pivot, juste pousser
// la moitie
func sortB(st ss.SortStacks, so *Solver) {
	pivot := st.B(so.lenSubStack()-1)		// value
	so.newSubStackLen = 0

	//st.PrintBoth()
/*	fmt.Println(pivot, right)
	st.PrintOps()
*/
	for st.FirstB() != pivot {
		st.PrintBoth()
		if st.FirstB() < pivot {
			pushA(st)
			so.newSubStackLen++
			so.decLenSubStack()
		} else {
			rotateB(st)
		}
	}
	//st.PrintBoth()
/*	for st.FirstB() != pivot {
		if st.FirstB() < pivot {
			pushA()
			new
		}
	}*/

	os.Exit(0)
}

/*
	last element is chosen as pivot first,
	if no inferior element is found
	the previous is picked
*/

func PivotA(st ss.SortStacks) int {
	var h func(int, ss.SortStacks) int
	h = func(i int, st ss.SortStacks) int {
		if (st.A(i)  == st.FirstA()) {
			return st.A(i)
		}
		for j := 0; i > j; j++ {
			if st.A(i) > st.A(j) {
				return st.A(i)
			}
		}
		return h(i-1, st)
	}
	return h(st.LenA()-1, st)
}

func log(p int, st ss.SortStacks) {
	fmt.Println(p)
	st.PrintBoth()
	st.PrintOps()
	os.Exit(0)
}

func n(f func(ss.SortStacks), n int, st ss.SortStacks) {
	for ; n > 0; n-- {
		f(st)
	}
}

func solve(st ss.SortStacks) {
	so := &Solver{
		lenSubStacks: make([]int, 1, 8),
	}
	p := PivotA(st)
	so.lenSubStacks[0] = st.LenA()


	//for st.Sorted() == false {
	for i := 0; i < so.lenSubStack() + so.newSubStackLen; i++ {
		if st.FirstA() <= p {
			pushB(st)
			so.decLenSubStack()
			so.newSubStackLen++
		} else if st.FirstA() > p {
			rotateA(st)
			so.rotations++
		}
	}
	// Every e elements such as e < p are now in B
	if st.LenSorted() > 0 {
		if so.lenSubStack() + st.LenSorted() < so.rotations  { // better rotate
			n(rotateA, so.lenSubStack() + st.LenSorted(), st)
		} else { // better revRotate
			n(revRotateA, so.rotations, st)
		}
	}

	so.lenSubStacks = append(so.lenSubStacks, so.newSubStackLen)
	log(p, st)
	//sortB(st, so)
	//			pivot = newPivot(solver, stacks)
	//}
}

/*func newPivot(stacks ss.SortStacks, solver *Solver) {
	if solver.minValIsSet
}*/


func main() {
	st := ss.New(parser.BuildStack(os.Args[1:]))
	solve(st)
	st.PrintOps()
	fmt.Println("not crashed")
}

// Used to find minimum value index in A to determine the new pivot
// which is A[minVal -1]

func indexOf(val int, a []int) int {
	for i, v := range a {
		if v == val {
			return i
		}
	}
	return -1
}
