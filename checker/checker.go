package main

import (
	"os"
	"fmt"
	"sort"
	"gitlab.com/louisportay/push_swap/parser"
)

func end_game(a []int, b []int) {
	if len(b) > 0 || sort.IntsAreSorted(a) == false {
		fmt.Println("KO")
	} else {
		fmt.Println("OK")
	}
}

func main() {
	a := parser.BuildStack(os.Args[1:])
	raw_m := parser.GetMoves()
	m := parser.CheckMoves(raw_m)
	a, b := play(a, m)
	end_game(a, b)
}
