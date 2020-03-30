package main

import (
	"fmt"
	"gitlab.com/louisportay/push_swap/stacks"
	"testing"
)

type indexer int

const (
	A indexer = iota + 0
	B
)

func TestPivot(t *testing.T) {
	var tests = []struct {
		subStacklen int
		st          stacks.Sorter
		ix          indexer
		want        int
	}{
		{5, stacks.NewSorterTest([]int{1, 3, 5, 2, 4}, []int{}), A, 3},
		{4, stacks.NewSorterTest([]int{1, 3, 5, 2, 4}, []int{}), A, 3},
		{3, stacks.NewSorterTest([]int{1, 3, 5, 2, 4}, []int{}), A, 3},
		{6, stacks.NewSorterTest([]int{1, 3, 5, 2, 4, 6}, []int{}), A, 4},
		{4, stacks.NewSorterTest([]int{1, 6, 5, 2, 4, 3}, []int{}), A, 5},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v - %v", tt.st.CopyA(), tt.subStacklen)
		t.Run(testname, func(t *testing.T) {
			var r int
			if tt.ix == A {
				r = Pivot(tt.st.A, tt.subStacklen)
			} else {
				r = Pivot(tt.st.B, tt.subStacklen)
			}
			if r != tt.want {
				t.Errorf("got %d, want %d", r, tt.want)
			}
		})
	}
}
