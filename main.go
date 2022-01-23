package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // nil
	curr := head

	for curr != nil {
		next := curr.Next // read next; only used for loop.

		// 2 switch operations
		curr.Next = prev // make next, the previous
		prev = curr      // make prev, the current

		curr = next // make curr, the next, for the following loop
	}
	return prev
}

/*
1 -> 2 -> 3 -> 4
step 0
curr = 1
----------------
step 1
next = 2
curr.next = nil
previous = 1
curr = 2
1 -> nil
----------------
step 2
next = 3
curr.next = 1
previous = 2
curr = 3
2 -> 1 -> nil
*/
