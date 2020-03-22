package sortstacks

import (
	"fmt"
	"log"
	"sort"
)

type Op func()

type sortStacks struct {
	a, b []int
	ops  []string
	f    map[string]Op
}

type SortStacks interface {
	SetOps([]string)
	AddOps([]string)
	Ops() []string
	Op(string) Op
	PrintOps()
	A() []int
	B() []int
	FirstA() int
	FirstB() int
	LastA() int
	LastB() int
	PrintA()
	PrintB()
	PrintBoth()
	IsASorted() bool

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
}

func (st *sortStacks) PrintA() {
	fmt.Printf("A: %v\n", st.a)
}

func (st *sortStacks) PrintB() {
	fmt.Printf("B: %v\n", st.b)
}

func (st *sortStacks) PrintBoth() {
	st.PrintA()
	st.PrintB()
}

func (st *sortStacks) A() []int {
	return st.a
}

func (st *sortStacks) B() []int {
	return st.b
}

func (st *sortStacks) AddOps(s []string) {
	st.ops = append(st.ops, s...)
}

func (st *sortStacks) Op(op string) Op {
	return st.f[op]
}

func (st *sortStacks) Ops() []string {
	return st.ops
}

func (st *sortStacks) SetOps(s []string) {
	for _, o := range s {
		if _, ok := st.f[o]; ok == false {
			log.Fatalln(fmt.Errorf("%v: not a valid operation", o))
		}
	}
	st.ops = s
}

func (st *sortStacks) PrintOps() {
	for _, v := range st.ops {
		fmt.Printf("%v\n", v)
	}
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

func New(i []int) SortStacks {
	st := &sortStacks{
		a: i,
		b: []int{},
	}
	st.helperNew()
	return st
}

// B list is empty and A is sorted in ascending order

func (st *sortStacks) IsASorted() bool {
	if len(st.b) > 0 || sort.IntsAreSorted(st.a) == false {
		return false
	} else {
		return true
	}
}
