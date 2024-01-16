package main

import "fmt"

type Stack []interface{}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(data interface{}) {
	*s = append(*s, data)
}

func (s *Stack) pop() interface{} {
	if s.isEmpty() {
		return nil
	} else {
		top := len(*s) - 1
		topData := (*s)[top]
		*s = (*s)[:top]
		return topData
	}
}

func main() {
	var s Stack
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Printf("%d poped from stack\n", s.pop())
}
