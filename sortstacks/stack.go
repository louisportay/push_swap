package sortstacks

import (
	"fmt"
	"log"
	"sort"
)

type sortStacks struct {
	a, b []int
	ops  []string
	f    map[string]func()
}

type SortStacks interface {
	// helperNew()
	SetOps([]string)
	CheckOps()
	GetFunc(op string) func()
	Play()
	IsOk() bool

	/*	SwapA()
		SwapB()
		SwapBoth()
		PushA()
		PushB()
		RotateA()
		RotateB()
		RotateBoth()
		RevRotateA()
		RevRotateB()
		RevRotateBoth()*/
}

func (st *sortStacks) CheckOps() {
	for _, o := range st.ops {
		if _, ok := st.f[o]; ok == false {
			log.Fatalln(fmt.Errorf("%v: not a valid operation", o))
		}
	}
}

func (st *sortStacks) GetFunc(s string) func() {
	return st.f[s]
}

func (st *sortStacks) SetOps(s []string) {
	st.ops = s
}

func (st *sortStacks) GetOps() []string {
	return st.ops
}

func (st *sortStacks) helperNew() {
	st.f = map[string]func(){
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

func (st *sortStacks) Play() {
	for _, o := range st.ops {
		fmt.Println(st.a, st.b) // VERBOSE
		st.f[o]()
	}
	fmt.Println(st.a, st.b) // VERBOSE
}

func (st *sortStacks) IsOk() bool {
	if len(st.b) > 0 || sort.IntsAreSorted(st.a) == false {
		return false
	} else {
		return true
	}
}
