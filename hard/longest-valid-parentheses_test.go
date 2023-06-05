package hard

import (
	"fmt"
	"testing"
)

var longestValidParenthesesCase = []string{
	"(()",
	")()())",
	"",
	"(((((((((())))))))))",
}

func TestLongestValidParentheses(t *testing.T) {
	for _, s := range longestValidParenthesesCase {
		fmt.Println(longestValidParentheses(s), " case:", s)
	}
}

func TestLongestValidParentheses2(t *testing.T) {
	for _, s := range longestValidParenthesesCase {
		fmt.Println(longestValidParentheses2(s), " case:", s)
	}
}
