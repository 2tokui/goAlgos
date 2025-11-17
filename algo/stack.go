// LIFO data structure
// stack objects on top of each other like a _tower of books_
// ppl standardized push() and pop() names
package algo

type Stack struct {
	container []any
}

func (s *Stack) Push(item any) {
	s.container = append(s.container, item)
}

func (s *Stack) Pop() any {
	if len(s.container) > 0 {
		idx := len(s.container)-1
		item := s.container[idx]
		s.container = s.container[0:idx]
		return item
	}
	return nil
}

func (s *Stack) Peek() any {
	if len(s.container) > 0 {
		return s.container[len(s.container)-1]
	}
	return nil
}

func (s *Stack) Empty() bool { return len(s.container) == 0}

func NewStack() *Stack {
	return &Stack{}
}

