package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	four := &ListNode{Val: 4, Next: nil}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}

	out := reverseList(one)
	want := 4
	if out.Val != want {
		t.Errorf("got %v, want %v", out, want)
	}
}
