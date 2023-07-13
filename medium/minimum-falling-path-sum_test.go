package medium

import (
	"fmt"
	"testing"
)

var case931 = []struct {
	c [][]int
	r int
}{
	{
		c: [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}},
		r: 13,
	},
	{
		c: [][]int{{-19, 57}, {-40, -5}},
		r: -59,
	},
	{
		c: [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}},
		r: 13,
	},
	{
		c: [][]int{{100, -42, -46, -41}, {31, 97, 10, -10}, {-58, -51, 82, 89}, {51, 81, 69, -51}},
		r: -36,
	},
}

func TestMinFallingPathSum(t *testing.T) {
	for _, v := range case931 {
		if res := minFallingPathSum(v.c); res != v.r {
			fmt.Println("case:", v.c, " res:", res, " should:", v.r)
			t.Fail()
		}
	}
}

func TestMinFallingPathSum1(t *testing.T) {
	for _, v := range case931 {
		if res := minFallingPathSum1(v.c); res != v.r {
			fmt.Println("case:", v.c, " res:", res, " should:", v.r)
			t.Fail()
		}
	}
}

func BenchmarkMinFallingPathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if res := minFallingPathSum([][]int{{100, -42, -46, -41}, {31, 97, 10, -10}, {-58, -51, 82, 89}, {51, 81, 69, -51}}); res != -36 {
			b.Fail()
		}
	}
}

func BenchmarkMinFallingPathSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if res := minFallingPathSum1([][]int{{100, -42, -46, -41}, {31, 97, 10, -10}, {-58, -51, 82, 89}, {51, 81, 69, -51}}); res != -36 {
			b.Fail()
		}
	}
}
