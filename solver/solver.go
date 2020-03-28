package main

import (
	"gitlab.com/louisportay/push_swap/parser"
	"gitlab.com/louisportay/push_swap/stacks"
	"os"
)

// OPTI: tester combien d'elements sont tries depuis le haut de A
// et adapter la strategie

// OPTI: regarder depuis la fin de A combien d'elements sont deja sorted et les ajouter a la liste
func main() {
	st := stacks.NewSorter(parser.BuildStack(os.Args[1:]))
	ss := New(st.LenA())
	sortA(st, ss)
	st.PrintOps()
}
