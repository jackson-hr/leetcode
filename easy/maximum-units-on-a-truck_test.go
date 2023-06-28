package easy

import (
	"log"
	"testing"
)

type case1710Struct struct {
	number [][]int
	size   int
	ans    int
}

var case1710 = []case1710Struct{
	{number: [][]int{{1, 3}, {2, 2}, {3, 1}}, size: 4, ans: 8},
	{number: [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, size: 10, ans: 91},
}

func TestMaximumUnits(t *testing.T) {
	for i := 0; i < len(case1710); i++ {
		var item = case1710[i]
		if res := maximumUnits(item.number, item.size); res != item.ans {
			log.Println("case:", item.number, " size:", item.size, " ans:", item.ans, " return:", res)
			t.Fail()
		}
	}
}
func TestMaximumUnits1(t *testing.T) {
	for i := 0; i < len(case1710); i++ {
		var item = case1710[i]
		if res := maximumUnits1(item.number, item.size); res != item.ans {
			log.Println("case:", item.number, " size:", item.size, " ans:", item.ans, " return:", res)
			t.Fail()
		}
	}
}

func BenchmarkMaximumUnits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if res := maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10); res != 91 {
			b.Fail()
		}
	}
}
func BenchmarkMaximumUnits1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if res := maximumUnits1([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10); res != 91 {
			b.Fail()
		}
	}
}
