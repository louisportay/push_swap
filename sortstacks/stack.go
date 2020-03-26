package sortstacks

import (
	"fmt"
	"log"
	"sort"
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

type SortStacks interface {
	SetOps([]string)
	AddOps([]string)
	Ops() []string
	Op(string) Op
	PrintOps()
	A(int) int
	B(int) int
	LenA() int
	LenB() int
	LenSorted() int
	FirstA() int
	FirstB() int
	LastA() int
	LastB() int
	PrintA()
	PrintB()
	PrintBoth()
	Sorted() bool

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

	RotateSorted()
	PushSorted()
	PrintSorted()
}

func reverseInts(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

/*
	'sorted' list is stored in ascending order
	'virtual A' list is 'sorted' + reversed 'A' list
	'B' is the only list that needs reversing
*/

func printStack(l []int, s string) {
	fmt.Printf("%v: %v\n", s, l)
}

func (st *sortStacks) PrintA() {
	printStack(st.virtualA(), "A")
}

func (st *sortStacks) PrintB() {
	tmp := make([]int, len(st.b))
	copy(tmp, st.b)
	reverseInts(tmp)
	printStack(tmp, "B")
}

func (st *sortStacks) PrintSorted() {
	printStack(st.sorted, "Sorted")

}

func (st *sortStacks) PrintBoth() {
	st.PrintA()
	st.PrintB()
}

func (st *sortStacks) A(i int) int {
	return st.a[len(st.a)-1-i]
}

func (st *sortStacks) B(i int) int {
	return st.b[len(st.b)-1-i]
}

func (st *sortStacks) LenA() int {
	return len(st.a)
}

func (st *sortStacks) LenB() int {
	return len(st.b)
}

func (st *sortStacks) LenSorted() int {
	return len(st.sorted)
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
	reverseInts(i)
	st := &sortStacks{
		a:      i,
		b:      make([]int, 0, len(i)),
		sorted: make([]int, 0, len(i)),
	}
	st.helperNew()
	return st
}

/*
	B list must be empty and
	A is first reversed and then concatenated with
	the 'sorted' special list
*/

//OPTI: this function is heavily called
// avoid allocing every time
func (st *sortStacks) virtualA() []int {
	tmp := make([]int, len(st.a))
	copy(tmp, st.a)
	reverseInts(tmp)
	tmp = append(tmp, st.sorted...)

	return tmp
}

func (st *sortStacks) Sorted() bool {
	if len(st.b) > 0 || sort.IntsAreSorted(st.virtualA()) == false {
		return false
	} else {
		return true
	}
}
