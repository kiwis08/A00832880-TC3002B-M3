package main

type Stack struct {
	elements []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top
}

func (s *Stack) Peek() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	return s.elements[len(s.elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Clear() {
	s.elements = []interface{}{}
}
