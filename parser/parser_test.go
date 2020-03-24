package parser

import (
	"testing"
	"fmt"
)

func eq(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBuildStack(t *testing.T) {
	var tests = []struct {
		args []string
		want []int
	}{
		{[]string{"1","2","3"}, []int{1,2,3}},
		{[]string{"1"}, []int{1}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.args)
		t.Run(testname, func(t *testing.T) {
			ret := BuildStack(tt.args)
			if eq(ret, tt.want) == false {
				t.Errorf("got: %v, want: %v", ret, tt.want)
			}
		})
	}
}
