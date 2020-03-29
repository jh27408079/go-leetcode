package main

import (
	"math/rand"
	"time"
)

type ListNode struct {
	Val int
	Next *ListNode
}

var arr []*ListNode

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
	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		arr = append(arr, initList(rand.Intn(10), rand.Intn(100), rand.Intn(10)))
	}
}

func printList(head *ListNode)  {

	for head != nil {
		println(head.Val)
		head = head.Next
	}
}

func main() {
	head := mergeKLists(arr)
	printList(head)
	println(333333)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return  nil
	}
	heap := []*ListNode{nil}
	var ret, p *ListNode
	for _, head := range lists {
		if head != nil {
			heap = append(heap, head)
			heap = swim(heap)
		}
	}
	ret = new(ListNode)
	p = ret
	for heap[1] != nil {
		p.Next = heap[1]
		p = p.Next
		heap[1] = heap[1].Next
		sink(heap)
	}
	ret = ret.Next
	return ret
}

func swim(heap []*ListNode) []*ListNode {
	i := len(heap) - 1
	if heap[i] == nil {
		return heap[:i - 1]
	}
	for i > 1 {
		if heap[i].Val < heap[i/2].Val {
			heap[i], heap[i/2] = heap[i/2], heap[i]
		}
		i /= 2
	}
	return heap
}

func sink(heap []*ListNode)  {
	i := 1
	l := len(heap) - 1

	for i <= l  {
		left := i * 2
		right := i * 2 + 1
		min := 0

		switch {
		case left > l:
			return
		case left == l:
			if heap[left] == nil {
				return
			} else {
				min = left
			}
		default:
			if heap[left] == nil && heap[right] == nil {
				return
			}
			if heap[left] == nil {
				min = right
			} else if heap[right] == nil {
				min = left
			} else if heap[left].Val > heap[right].Val {
				min = right
			} else {
				min = left
			}
		}
		if heap[i] == nil || heap[i].Val > heap[min].Val {
			heap[i], heap[min] = heap[min], heap[i]
			i = min
			continue
		}
		return
	}
}


