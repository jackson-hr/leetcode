package easy

import (
	"fmt"
	"testing"
)

var caseInterview1704 = []struct {
	c []int
	r int
}{
	{c: []int{3, 0, 1}, r: 2},
	{c: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, r: 8},
	{c: []int{0}, r: 1},
}

func TestMissingNumber(t *testing.T) {
	for _, s := range caseInterview1704 {
		if res := missingNumber(s.c); res != s.r {
			fmt.Println("case:", s.c, " res:", res, " should:", s.r)
			t.Fail()
		}
	}
}
