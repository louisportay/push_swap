package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func BuildStack(args []string) (s []int) {
	if len(args) == 0 {
		os.Exit(0)
	}
	cache := make(map[int]bool)
	for _, v := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		} else if _, ok := cache[val]; ok == true {
			log.Fatalln(fmt.Errorf("%v: duplicated value", val))
		}
		cache[val] = true
		s = append(s, val)
	}
	return s
}

func ReadOps() (m []string) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() == true {
		o := strings.Trim(s.Text(), " 	")
		if len(o) > 0 {
			m = append(m, o)
		}
	}
	if e := s.Err(); e != nil {
		log.Fatalln(e)
	}
	return m
}
