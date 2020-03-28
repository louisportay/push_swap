package stacks

import (
	"fmt"
	"log"
)

type Op func()

/*
	Virtually, the 'sorted' list is contained in A but for implementation
	optimizations, it has a dedicated list
*/

type sortStacks struct {
	a, b   []int
	sorted []int
	ops    []string
	f      map[string]Op
}

type Sorter interface {
	AddOps([]string)
	Ops() []string
	Op(string) Op
	PrintOps()

	A(int) int
	B(int) int
	LenA() int
	LenB() int
	CopyA() []int
	CopyB() []int
	LastA() int
	LastB() int
	PrintA()
	PrintB()

	Sorted(int) int
	LenSorted() int
	CopySorted() []int
	PrintSorted()

	SwapA()
	SwapB()
	SwapBoth()
	PushA()
	PushB()
	RotateA()
	RotateB()
	RotateBoth()
	RevRotateA()
	RevRotateB()
	RevRotateBoth()

	PushASorted()
	PushBSorted()
}

func reverseInts(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func (st *sortStacks) checkOps(s []string) {
	for _, o := range s {
		if _, ok := st.f[o]; ok == false {
			log.Fatalln(fmt.Errorf("%v: not a valid operation", o))
		}
	}
}

func (st *sortStacks) AddOps(s []string) {
	st.checkOps(s)
	st.ops = append(st.ops, s...)
}

func (st *sortStacks) Ops() []string {
	return st.ops
}

func (st *sortStacks) helperNew() {
	st.f = map[string]Op{
		"sa":  st.SwapA,
		"sb":  st.SwapB,
		"ss":  st.SwapBoth,
		"pa":  st.PushA,
		"pb":  st.PushB,
		"ra":  st.RotateA,
		"rb":  st.RotateB,
		"rr":  st.RotateBoth,
		"rra": st.RevRotateA,
		"rrb": st.RevRotateB,
		"rrr": st.RevRotateBoth,
	}
}

/*
	intended for testing purposes only
	'B' is always 'nil' in normal spawning conditions
*/
func NewSorterTest(A []int, B []int) Sorter {
	reverseInts(A)
	reverseInts(B)
	st := &sortStacks{
		a:      A,
		b:      B,
		sorted: make([]int, 0, len(A)),
	}
	st.helperNew()
	return st
}

func NewSorter(i []int) Sorter {
	reverseInts(i)
	st := &sortStacks{
		a:      i,
		b:      make([]int, 0, len(i)),
		sorted: make([]int, 0, len(i)),
	}
	st.helperNew()
	return st
}
