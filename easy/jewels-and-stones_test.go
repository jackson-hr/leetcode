package easy

import (
	"fmt"
	"testing"
)

var case771 = []struct {
	j string
	s string
	r int
}{
	{j: "aA", s: "aAAbbbb", r: 3},
	{j: "z", s: "ZZ", r: 0},
}

func TestNumJewelsInStones(t *testing.T) {
	for i := range case771 {
		if res := numJewelsInStones(case771[i].j, case771[i].s); res != case771[i].r {
			fmt.Println("case ", case771[i], "res ", res)
			t.Fail()
		}
	}
}
