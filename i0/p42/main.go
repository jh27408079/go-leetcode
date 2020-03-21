package main

func main() {
	height := []int{4,2,0,3,2,5}
	println(trap(height))
}

func trap(height []int) int {
	s := 0
	st := newStack()
	l := len(height)
	st.Push(0)
	for i := 1; i < l; i++ {
		for !st.IsEmpty() && height[i] > height[st.Top.Item] {
			_, top := st.Pop()
			if st.IsEmpty() {
				break
			}
			distance := i - st.Top.Item - 1
			boundedHeight := height[i]
			if height[st.Top.Item] < boundedHeight {
				boundedHeight = height[st.Top.Item]
			}
			boundedHeight -= height[top]
			s += boundedHeight * distance
		}
		println(s)
		st.Push(i)
	}
	return s
}

type node struct {
	Item int
	Next *node
}

type stack struct {
	Top *node
}

func newStack() stack {
	return stack{}
}

func (s *stack)IsEmpty() bool {
	return s.Top == nil
}

func (s *stack)Push(item int)  {
	n := node{
		Item: item,
		Next: s.Top,
	}
	s.Top = &n
}

func (s *stack) Pop() (ok bool, item int)  {
	if s.IsEmpty() {
		return false, 0
	}
	ok = true
	item = s.Top.Item
	s.Top = s.Top.Next
	return
}
