package main

import (
	"strings"
)

func main() {
	println(simplifyPath("/home/../.././//aaa/bbb/../ccc"))
}

func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	s := newStack()
	for _, item := range arr {
		if item == "." || item == "" {
			continue
		} else if item == ".." {
			s.Pop()
		} else {
			s.Push(item)
		}
	}
	return s.ToPath()
}


type stack []string

func newStack() stack {
	return stack(make([]string, 0))
}
func (s *stack) Push(item string) {
	*s = append(*s, item)
	return
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Pop() (ok bool, item string) {
	l := len(*s)
	if l == 0 {
		return false, ""
	}
	ok = true
	item = (*s)[l-1]
	*s = (*s)[:l-1]
	return
}

func (s *stack) ToPath() string {
	ret := "/"
	ret += strings.Join(*s, "/")
	return ret
}