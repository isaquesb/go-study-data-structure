package main

type DequeueStrategy interface {
	Dequeue(q *Queue) NodeContent
}

type StrategyAbstract struct {
}

func (f *StrategyAbstract) setHead(m *Node, q *Queue) {
	q.iterator.head = m
	if q.iterator.head != nil {
		q.iterator.head.prev = nil
	} else {
		f.setTail(nil, q)
	}
}

func (f *StrategyAbstract) setTail(m *Node, q *Queue) {
	q.iterator.tail = m
	if q.iterator.tail != nil {
		q.iterator.tail.next = nil
	}
}
