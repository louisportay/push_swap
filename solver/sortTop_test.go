package main

import (
	"testing"
	"fmt"
	s "gitlab.com/louisportay/push_swap/sortstacks"
)


func eqIntSlices(got []int, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	for i := 0; i <  len(got); i++ {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func eqStringSlices(got []string, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := 0; i <  len(got); i++ {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func TestSortTopA(t *testing.T) {
	var tests = []struct {
		subStacklen int
		st s.SortStacks
		wantedA []int
		wantedB []int
		wantedSorted []int
		wantedOps []string
	}{
		{1, s.NewTest([]int{1,3,8,19},[]int{2}),[]int{3,8,19},[]int{2},[]int{1},[]string{"ra"}},
		{2, s.NewTest([]int{1,-1,3,8,19},[]int{2}),[]int{3,8,19},[]int{2},[]int{-1,1},[]string{"sa","ra","ra"}},
		{2, s.NewTest([]int{-1,1,3,8,19},[]int{2}),[]int{3,8,19},[]int{2},[]int{-1,1},[]string{"ra","ra"}},
		{3, s.NewTest([]int{3,4,5,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{3,4,5},[]string{"ra","ra","ra"}},
		{3, s.NewTest([]int{3,5,4,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{3,4,5},[]string{"ra","sa","ra","ra"}},
		{3, s.NewTest([]int{4,3,5,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{3,4,5},[]string{"sa","ra","ra","ra"}},
		{3, s.NewTest([]int{4,5,3,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{3,4,5},[]string{"pb","sa","ra","pa","ra","ra"}},
		{3, s.NewTest([]int{5,4,3,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{3,4,5},[]string{"pb","sa","ra","ra","pa", "ra"}},
		{3, s.NewTest([]int{5,3,4,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{3,4,5},[]string{"sa","ra","sa","ra","ra"}},
		{4, s.NewTest([]int{2,3,4,5,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"ra","ra","ra","ra"}},
		{4, s.NewTest([]int{2,3,5,4,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"ra","ra","sa","ra","ra"}},
		{4, s.NewTest([]int{2,4,5,3,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"ra","pb","sa","ra","pa","ra","ra"}},
		{4, s.NewTest([]int{2,4,3,5,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"ra","sa","ra","ra","ra"}},
		{4, s.NewTest([]int{2,5,4,3,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"ra","pb","sa","ra","ra","pa","ra"}},
		{4, s.NewTest([]int{2,5,3,4,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"ra","sa","ra","sa","ra","ra"}},
		{4, s.NewTest([]int{3,2,4,5,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"sa","ra","ra","ra","ra"}},
		{4, s.NewTest([]int{3,2,5,4,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"sa","ra","ra","sa","ra","ra"}},
		{4, s.NewTest([]int{3,4,5,2,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"pb","pb","sa","ra","sb","pa","ra","pa","ra","ra"}},
		{4, s.NewTest([]int{3,4,2,5,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"pb","sa","ra","pa","ra","ra","ra"}},
		{4, s.NewTest([]int{3,5,4,2,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"pb","pb","sa","ra","sb","pa","ra","ra","pa","ra"}},
		{4, s.NewTest([]int{3,5,2,4,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"pb","sa","ra","pa","ra","sa","ra","ra"}},
		{4, s.NewTest([]int{4,3,2,5,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"pb","sa","ra","ra","pa","ra","ra"}},
		{4, s.NewTest([]int{4,3,5,2,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"pb","pb","sa","ra","pa","ra","pa","ra","ra"}},
		{4, s.NewTest([]int{4,2,5,3,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"sa","ra","pb","sa","ra","pa","ra","ra"}},
		{4, s.NewTest([]int{4,2,3,5,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"sa","ra","sa","ra","ra","ra"}},
		{4, s.NewTest([]int{4,5,2,3,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"pb","sa","ra","sa","ra","pa","ra","ra"}},
		{4, s.NewTest([]int{4,5,3,2,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"pb","pb","sa","ra","ra","sb","pa","ra","pa","ra"}},
		{4, s.NewTest([]int{5,3,4,2,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"pb","pb","sa","ra","pa","ra","ra","pa","ra"}},
		{4, s.NewTest([]int{5,3,2,4,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"pb","sa","ra","ra","ra","pa","ra"}},
		{4, s.NewTest([]int{5,4,2,3,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"pb","sa","ra","sa","ra","ra","pa","ra"}},
		{4, s.NewTest([]int{5,4,3,2,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"pb","pb","sa","ra","ra","pa","ra","pa","ra"}},
		{4, s.NewTest([]int{5,2,4,3,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"sa","ra","pb","sa","ra","ra","pa","ra"}},
		{4, s.NewTest([]int{5,2,3,4,21,8,19},[]int{10}),[]int{21,8,19},[]int{10},[]int{2,3,4,5},[]string{"sa","ra","sa","ra","sa","ra","ra"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.st.CopyA())
		t.Run(testname, func(t *testing.T) {
			sortTopA(tt.subStacklen, tt.st)
			if eqIntSlices(tt.st.CopyA(), tt.wantedA) == false {
				t.Errorf("A got: %v, want: %v", tt.st.CopyA(), tt.wantedA)
			}
			if eqIntSlices(tt.st.CopyB(), tt.wantedB) == false {
				t.Errorf("B got: %v, want: %v", tt.st.CopyB(), tt.wantedB)
			}
			if eqIntSlices(tt.st.CopySorted(), tt.wantedSorted) == false {
				t.Errorf("Sorted got: %v, want: %v", tt.st.CopySorted(), tt.wantedSorted)
			}
			if eqStringSlices(tt.st.Ops(), tt.wantedOps) == false {
				t.Errorf("Ops got: %v, want: %v", tt.st.Ops(), tt.wantedOps)
			}
		})
	}
}
