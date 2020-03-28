package stacks

import "fmt"

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

func (st *sortStacks) PrintOps() {
	for _, v := range st.ops {
		fmt.Printf("%v\n", v)
	}
}
