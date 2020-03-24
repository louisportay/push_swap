package main

import (
	"fmt"
	"gitlab.com/louisportay/push_swap/parser"
	ss "gitlab.com/louisportay/push_swap/sortstacks"
	"os"
)

func displayScore(ok bool) {
	if ok {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

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

func play(s ss.SortStacks, v bool) {
	for _, o := range s.Ops() {
		if v == true {
			s.PrintBoth()
			fmt.Printf("--- %v ---\n", o)
		}
		s.Op(o)()
	}
	if v == true {
		s.PrintBoth()
	}
}

func main() {
	verbose := setFlags()
	s := ss.New(parser.BuildStack(os.Args[1:]))
	s.SetOps(parser.ReadOps())
	play(s, verbose)
	displayScore(s.Sorted())
}
