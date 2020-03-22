package main

import (
	"fmt"
	"gitlab.com/louisportay/push_swap/parser"
	st "gitlab.com/louisportay/push_swap/sortstacks"
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
		fmt.Println("https://cdn.intra.42.fr/pdf/pdf/9430/fr.subject.pdf")
		os.Exit(0)
	} else if os.Args[1] == "-v" {
		os.Args = os.Args[1:]
		return true
	}
	return false
}

func play(s st.SortStacks, v bool) {
	for _, o := range s.Ops() {
		if v == true {
			s.PrintBoth()
		}
		s.Op(o)()
	}
	if v == true {
		s.PrintBoth()
	}
}

func main() {
	verbose := setFlags()
	s := st.New(parser.BuildStack(os.Args[1:]))
	s.SetOps(parser.ReadOps())
	play(s, verbose)
	displayScore(s.IsASorted())
}
