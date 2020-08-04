package test

import "testing"

type charcountdata struct {
	s     string
	count int
}

func TestCharCount(t *testing.T) {
	input := []charcountdata{{"rrm", 3}, {" rrm", 3}, {"r rm", 3}, {"rrm ", 3}, {"hello! i am waiting", 16}, {"1 2 ", 2}}

	for _, in := range input {
		if out := CharCount(in.s); out != in.count {
			t.Errorf("String : %q\tCount Expected :%d\t Count :%d\n", in.s, in.count, out)
		}
	}
}
