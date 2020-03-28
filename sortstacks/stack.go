package sortstacks

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

type SortStacks interface {
	SetOps([]string)
	AddOps([]string)
	Ops() []string
	Op(string) Op
	PrintOps()
	A(int) int
	B(int) int
	Sorted(int) int
	LenA() int
	LenB() int
	LenSorted() int
	CopyA() []int
	CopyB() []int
	CopySorted() []int
	FirstA() int
	FirstB() int
	LastA() int
	LastB() int
	PrintA()
	PrintB()

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
	'virtual A' list is made of 'sorted' + reversed 'A' lists
	'B' is the only list that needs explicit reversing
*/

func printStack(l []int, s string) {
	fmt.Printf("%v: %v\n", s, l)
}

func (st *sortStacks) PrintA() {
	printStack(st.virtualA(), "A")
}

func (st *sortStacks) PrintB() {
	printStack(st.CopyB(), "B")
}

func (st *sortStacks) PrintSorted() {
	printStack(st.sorted, "Sorted")

}

func (st *sortStacks) A(i int) int {
	return st.a[len(st.a)-1-i]
}

func (st *sortStacks) B(i int) int {
	return st.b[len(st.b)-1-i]
}

func (st *sortStacks) Sorted(i int) int {
	return st.sorted[i]
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

func (st *sortStacks) checkOps(s []string) {
	for _, o := range s {
		if _, ok := st.f[o]; ok == false {
			log.Fatalln(fmt.Errorf("%v: not a valid operation", o))
		}
	}
}

func (st *sortStacks) SetOps(s []string) {
	st.checkOps(s)
	st.ops = s
}

func (st *sortStacks) AddOps(s []string) {
	st.checkOps(s)
	st.ops = append(st.ops, s...)
}

func (st *sortStacks) Op(op string) Op {
	return st.f[op]
}

func (st *sortStacks) Ops() []string {
	return st.ops
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

/*
	intended for testing purposes only
	'B' is always 'nil' in normal conditions
*/
func NewTest(A []int, B []int) SortStacks {
	reverseInts(A)
	reverseInts(B)
	st := &sortStacks{
		a: A,
		b: B,
		sorted: make([]int, 0, len(A)),
	}
	st.helperNew()
	return st
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

func (st *sortStacks) virtualA() []int {
	tmp := make([]int, len(st.a))
	copy(tmp, st.a)
	reverseInts(tmp)
	tmp = append(tmp, st.sorted...)

	return tmp
}

func (st *sortStacks) CopyA() []int {
	cp := make([]int, len(st.a))
	copy(cp, st.a)
	reverseInts(cp)

	return cp
}

func (st *sortStacks) CopyB() []int {
	cp := make([]int, len(st.b))
	copy(cp, st.b)
	reverseInts(cp)

	return cp
}

func (st *sortStacks) CopySorted() []int {
	cp := make([]int, len(st.b))
	copy(cp, st.b)

	return cp
}
