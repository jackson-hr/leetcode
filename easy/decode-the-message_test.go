package easy

import (
	"fmt"
	"testing"
)

var case2325 = []struct{ k, m, r string }{
	{k: "the quick brown fox jumps over the lazy dog", m: "vkbs bs t suepuv", r: "this is a secret"},
	{k: "eljuxhpwnyrdgtqkviszcfmabo", m: "zwx hnfx lqantp mnoeius ycgk vcnjrdb", r: "the five boxing wizards jump quickly"},
}

func TestDecodeMessage(t *testing.T) {
	for _, s := range case2325 {
		if res := decodeMessage(s.k, s.m); res != s.r {
			fmt.Println("case:", s, " res:", res)
			t.Fail()
		}
	}
}
