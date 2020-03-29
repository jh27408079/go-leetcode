package main

type ListNode struct {
	Val int
	Next *ListNode
}


func initList(l, start, step int) *ListNode {
	var head, p *ListNode
	head = new(ListNode)
	p = head
	p.Val = start
	for i := 1; i < l; i++ {
		p.Next = new(ListNode)
		p = p.Next
		p.Val = start + i * step
	}
	return head
}

func init()  {

}

func printList(head *ListNode)  {

	for head != nil {
		println(head.Val)
		head = head.Next
	}
}

func main() {
	p := initList(4, 5, 3)
	printList(p)
	println()
	head := reverseKGroup(p, 3)
	printList(head)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var ret, p, q, rest *ListNode
	var stack []*ListNode
	ret = new(ListNode)
	p = ret
	for  {
		flag := false
		rest = head
		for i := 0; i < k; i++ {

			if head != nil {
				stack = push(stack, head)
				head = head.Next
			} else {
				flag = true
				break
			}
		}
		if flag {
			break
		}
		for len(stack) != 0 {
			stack, q = pop(stack)
			p.Next = q
			p = p.Next
		}
	}
	p.Next = rest
	ret = ret.Next
	return ret
}

func pop(stack []*ListNode) (s []*ListNode, node *ListNode) {
	l := len(stack)
	if len(stack) == 0 {
		return stack, nil
	} else {
		return stack[:l - 1], stack[l - 1]
	}
}

func push(stack []*ListNode, node *ListNode) []*ListNode {
	stack = append(stack, node)
	return stack
}