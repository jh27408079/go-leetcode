package main

// Constants to be used throughout the example

func main() {
	println(isValid("(([]{}))"))
}

func isValid(s string) bool {
	arr := []rune(s)
	match := make(map[rune]rune)
	match[')'] = '('
	match[']'] = '['
	match['}'] = '{'
	if len(arr) % 2 != 0 {
		return false
	}
	st := newStack()
	for _, item := range arr {
		if item == '(' || item == '[' || item == '{' {
			st.Push(item)
		} else if item == ')' || item == ']' || item == '}' {
			ok, temp := st.Pop()
			if !ok || match[item] != temp {
				return false
			}
		} else {
			return false
		}
	}
	return st.IsEmpty()
}

type node struct {
	Item rune
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

func (s *stack)Push(item rune)  {
	n := node{
		Item: item,
		Next: s.Top,
	}
	s.Top = &n
}

func (s *stack) Pop() (ok bool, item rune)  {
	if s.IsEmpty() {
		return false, 0
	}
	ok = true
	item = s.Top.Item
	s.Top = s.Top.Next
	return
}

//type stack []rune
//
//
//func newStack() stack {
//	return stack(make([]rune, 0))
//}
//func (s *stack) Push(item rune) {
//	*s = append(*s, item)
//	return
//}
//
//func (s *stack) IsEmpty() bool {
//	return len(*s) == 0
//}
//
//func (s *stack) Pop() (ok bool, item rune) {
//	l := len(*s)
//	if l == 0 {
//		return false, 0
//	}
//	ok = true
//	item = (*s)[l-1]
//	*s = (*s)[:l-1]
//	return
//}