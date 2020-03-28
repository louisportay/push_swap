package main

import (
	"fmt"
	"gitlab.com/louisportay/push_swap/parser"
	"gitlab.com/louisportay/push_swap/stacks"
	"os"
	"sort"
)

func setFlags() bool {
	if len(os.Args) == 1 {
		os.Exit(0)
	} else if os.Args[1] == "-h" {
		fmt.Println("")
		os.Exit(0)
	} else if os.Args[1] == "-v" {
		os.Args = os.Args[1:]
		return true
	}
	return false
}

func play(st stacks.Sorter, v bool) {
	for _, o := range st.Ops() {
		if v == true {
			st.PrintA()
			st.PrintB()
			fmt.Printf("--- %v ---\n", o)
		}
		st.Op(o)()
	}
	if v == true {
		st.PrintA()
		st.PrintB()
	}
}

func main() {
	verbose := setFlags()
	st := stacks.NewSorter(parser.BuildStack(os.Args[1:]))
	st.AddOps(parser.ReadOps())
	play(st, verbose)
	if st.LenB() == 0 && sort.IntsAreSorted(st.CopyA()) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
