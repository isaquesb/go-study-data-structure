package main

type Queue struct {
	pipe
}

func (q *Queue) Enqueue(c NodeContent) {
	q.addLast(c)
}

func (q *Queue) Dequeue() NodeContent {
	strategy := Fifo{}
	return strategy.pull(q)
}
