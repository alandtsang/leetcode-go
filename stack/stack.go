package stack

import (
	"container/list"
	"fmt"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New() // generate list
	return &Stack{list}
}

func (s *Stack) Push(v interface{}) {
	s.list.PushBack(v)
}

func (s *Stack) Pop() interface{} {
	e := s.list.Back()
	if e == nil {
		return nil
	}
	s.list.Remove(e)
	return e.Value
}

func (s *Stack) Peek() interface{} {
	e := s.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (s *Stack) Travel() {
	for e := s.list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}
