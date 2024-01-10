package easy

import (
	"log"
	"testing"
)

var case2696 = []struct {
	c string
	r int
}{
	{c: "ABFCACDB", r: 2},
	{c: "ACBBD", r: 5},
}

func TestMinLength(t *testing.T) {
	for _, item := range case2696 {
		res := minLength(item.c)
		log.Println("c:", item.c, " r:", item.r, " res:", res)
		if res != item.r {
			t.Fail()
		}
	}
}
