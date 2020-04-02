package oldflag

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	_ = t
	fs := NewFlagSet("", ExitOnError)
	a := fs.Bool('a', false, "aaa")
	c := fs.Int('c', 2, "aaccaccaaccaca")
	fs.Parse([]string{"-ac", "73", "-b", "--help", "b"})
	fmt.Println(*a, *c, fs.Args())
}
