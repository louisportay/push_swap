package parser

import (
	"os"
	"strconv"
	"strings"
	"bufio"
	"fmt"
)

func abort(e error) {
	fmt.Println(e)
	os.Exit(-1)
}

func BuildStack(args []string) (s []int) {
	if len(args) == 0 {
		os.Exit(0)
	}
	cache := make(map[int]bool)
	for _, v  := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			abort(err)
		} else if _, ok := cache[val]; ok == true {
			abort(fmt.Errorf("%v: duplicated value", val))
		}
		cache[val] = true
		s = append(s, val)
	}
	return s
}

func GetMoves() (m[]string) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() == true {
		m= append(m, strings.Trim(s.Text(), " 	"))
	}
	if e := s.Err(); e != nil {
		abort(e)
	}
	return m
}

func CheckMoves(moves []string) (intMoves []int) {
	allowedMoves := map[string]int {
		"sa": 1,
		"sb": 2,
		"ss": 3,
		"pa": 4,
		"pb": 5,
		"ra": 6,
		"rb": 7,
		"rr": 8,
		"rra": 9,
		"rrb": 10,
		"rrr": 11,
	}
	specialMoves := map[string]bool {
		"ss": true,
		"rr": true,
		"rrr": true,
	}
	for _, m := range moves {
		if _, ok := allowedMoves[m]; ok == false {
			abort(fmt.Errorf("%v: not a valid move", m))
		}
		if _, ok := specialMoves[m]; ok == true {
			intMoves = append(intMoves, allowedMoves[m] - 2)
			intMoves = append(intMoves, allowedMoves[m] - 1)
		} else {
			intMoves = append(intMoves, allowedMoves[m])
		}
	}
	return
}
