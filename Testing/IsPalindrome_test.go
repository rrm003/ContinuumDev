package test

import "testing"

type datapalindrome struct {
	data   string
	output bool
}

func TestIsPalindrome(t *testing.T) {
	input := []datapalindrome{{"rrm", false}, {"a", true}, {"aca", true}, {"aa", true}, {"aaabbbccc", false}}
	for _, in := range input {
		if out := IsPalindrome(in.data); out != in.output {
			t.Fatalf("Input :%q\t Output :%t Expected :%t\n", in.data, out, in.output)
		}
	}
}

func TestSpaceIsPalindrome(t *testing.T) {
	input := []datapalindrome{{"rr m", false}, {"a ", true}, {"a ca", true}, {" aa", true}, {" a aab bbc cc ", false}}
	for _, in := range input {
		if out := IsPalindrome(in.data); out != in.output {
			t.Fatalf("Input :%q\t Output :%t Expected :%t\n", in.data, out, in.output)
		}
	}
}

func TestCaseIsPalindrome(t *testing.T) {
	input := []datapalindrome{{"rR m", false}, {"a ", true}, {"a cA", true}, {" aA", true}, {" a aAb bbc Cc ", false}}
	for _, in := range input {
		if out := IsPalindrome(in.data); out != in.output {
			t.Fatalf("Input :%q\t Output :%t Expected :%t\n", in.data, out, in.output)
		}
	}
}
