package main

import "fmt"

func swap_a(a []int, b []int) ([]int, []int) {
	if len(a) >= 2 {
		a[0], a[1] = a[1], a[0]
		_ = b
	}
	return a, b
}

func swap_b(a []int, b []int) ([]int, []int) {
	if len(b) >= 2 {
		b[0], b[1] = b[1], b[0]
		_ = a
	}
	return a, b
}


func push_a(a []int, b []int) ([]int, []int) {
	if len(b) > 0 {
		var x int
		x, b = b[0], b[1:]
		a = append([]int{x}, a...)
	}
	return a, b
}

func push_b(a []int, b []int) ([]int, []int) {
	if len(a) > 0 {
		var x int
		x, a = a[0], a[1:]
		b = append([]int{x}, b...)
	}
	return a, b
}

func rotate_a(a []int, b []int) ([]int, []int) {
	if len(a) > 1 {
		var x int
		x, a = a[0], a[1:]
		a = append(a, x)
		_ = b
	}
	return a, b
}

func rotate_b(a []int, b []int) ([]int, []int) {
	if len(b) > 1 {
		var x int
		x, b = b[0], b[1:]
		b = append(b, x)
		_ = a
	}
	return a, b
}

func  rev_rotate_a(a []int, b []int) ([]int, []int) {
	if len(a) > 1 {
		var x int
		x, a = a[len(a)-1], a[:len(a)-1]
		a = append([]int{x}, a...)
		_ = b
	}
	return a, b
}

func  rev_rotate_b(a []int, b []int) ([]int, []int) {
	if len(b) > 1 {
		var x int
		x, b = b[len(b)-1], b[:len(b)-1]
		b = append([]int{x}, b...)
		_ = a
	}
	return a, b
}

func play(a []int, moves []int) []int {
	b := []int{}
	turn := map[int]func([]int, []int) ([]int, []int) {
		1: swap_a,
		2: swap_b,
		4: push_a,
		5: push_b,
		6: rotate_a,
		7: rotate_b,
		9: rev_rotate_a,
		10: rev_rotate_b,
	}
	for _, v := range moves {
		fmt.Println(a, b)//
		a, b = turn[v](a, b)
	}
	fmt.Println(a, b)//
	return a
}
