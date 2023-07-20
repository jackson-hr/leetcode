package easy

import (
	"fmt"
	"testing"
)

var case415 = []struct {
	num1 string
	num2 string
}{
	{"11", "123"},
	{"456", "77"},
	{"456", "777"},
	{"0", "0"},
}

func TestAddStrings(t *testing.T) {
	for _, s := range case415 {
		fmt.Println(addStrings(s.num1, s.num2), " num1:", s.num1, " num2:", s.num2)
	}
}
