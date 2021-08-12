package add_two_numbers

type ListNode struct {
	Val int 
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	}
	
	sum := l1.Val + l2.Val 
	next := addTwoNumbers(l1.Next, l2.Next)
	if sum >= 10 {
		//carry := 1
		//sum -= 10
		carry := sum / 10
		sum %= 10
		next = addTwoNumbers(next, &ListNode{Val: carry})
	}
	return &ListNode{Val: sum, Next:next}
}