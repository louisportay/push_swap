package main

import (
	"os"
	"gitlab.com/louisportay/push_swap/parser"
	s "gitlab.com/louisportay/push_swap/sortstacks"
)

// OPTI: tester cobien d'elements sont tries depuis le haut de A
// et adapter la strategie

func main() {
	st := s.New(parser.BuildStack(os.Args[1:]))
	ss := New(st.LenA())
	sortA(st, ss)
	st.PrintOps()
}
