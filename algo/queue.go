package algo

type Queue struct{
	Head any
	Tail any
	Collection []any
}

func NewQueue() *Queue {
	return &Queue{}
}

func (s *Queue) Enqueue(thing any) {
	s.Collection = append(s.Collection, thing)
	s.Tail = thing
	s.Head = s.Collection[0]
}

func (s *Queue) Dequeue() {
	if len(s.Collection) > 1 {
		s.Collection = s.Collection[1:]
		s.Head = s.Collection[0]
	} else {
		s.Head = nil
		s.Tail = nil
		s.Collection = []any{}
	}
}

func (s *Queue) Peek() any {
	return s.Head
}
