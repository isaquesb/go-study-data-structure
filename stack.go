package main

type Stack struct {
	pipe
}

func (s *Stack) Push(c NodeContent) {
	s.addLast(c)
}

func (s *Stack) Pop() NodeContent {
	strategy := Lifo{}
	return strategy.pull(s)
}
