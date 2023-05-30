package easy

import (
	"fmt"
	"testing"
)

var lastWords = []string{
	"Hello World",
	"   fly me   to   the moon  ",
	"luffy is still joyboy",
}

func TestLengthOfLastWord(t *testing.T) {
	for _, word := range lastWords {
		fmt.Println(lengthOfLastWord(word), " str:", word)
	}
}
