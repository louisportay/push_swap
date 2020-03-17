package main

import (
	"os"
	"fmt"
	"strings"
	//"reflect"
	"strconv"
	//"errors"
	"bufio"
)

func abort(e error) {
	fmt.Println(e)
	os.Exit(-1)
}

func buildStack(args []string) (s []int) {
	cache := make(map[int]bool)
	for _, v  := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			abort(err)
		} else if _, ok := cache[val]; ok == true {
			abort(fmt.Errorf("%v: duplicated value", val))
		}
		fmt.Println(cache)
		cache[val] = true
		s = append(s, val)
	}
	return s
}

func getMoves() (m[]string) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() == true {
		m= append(m, strings.Trim(s.Text(), " 	"))
	}
	if e := s.Err(); e != nil {
		abort(e)
	}
	return m
}

func checkMoves(moves []string) {
	allowedMoves := map[string]bool {
		"sa": true,
		"sb": true,
		"ss": true,
		"pa": true,
		"pb": true,
		"ra": true,
		"rb": true,
		"rr": true,
		"rra": true,
		"rrb": true,
		"rrr": true,
	}
	for _, m := range moves {
		if _, ok := allowedMoves[m]; ok == false {
			abort(fmt.Errorf("%v: not a valid move", m))
		}
	}
}

func main() {
	st := buildStack(os.Args[1:])
	m := getMoves()
	checkMoves(m)
}
