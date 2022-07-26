package stack

import (
	//	"container/list"
	"errors"
)

type Stack struct {
	Top  int
	Data []rune
	//stack *list.List
}

func (s *Stack) Push(Elemen rune) {
	s.Top += 1
	s.Data = append(s.Data, Elemen)
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) Pop() (rune, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack underflow")
	} else {
		poppedValue := s.Data[s.Top]
		s.Top -= 1
		s.Data = s.Data[:len(s.Data)-1]
		return poppedValue, nil
	}
}

func (s *Stack) Peek() (rune, error) {
	if s.Top == -1 {
		return 0, errors.New("stack is empty")
	} else {
		return s.Data[s.Top], nil
	}
}
