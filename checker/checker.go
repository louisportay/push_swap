package main

import (
	"os"
	"fmt"
	"gitlab.com/louisportay/push_swap/parser"
	st "gitlab.com/louisportay/push_swap/sortstacks"
)

func main() {
	s := st.New(parser.BuildStack(os.Args[1:]))
	s.SetOps(parser.GetOps())
	s.CheckOps()
	s.Play()
	if s.IsOk() == true {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
