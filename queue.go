package main

type Queue struct {
	iterator *Iterator
}

func (q *Queue) GetIterator() *Iterator {
	if q.iterator == nil {
		q.iterator = NewIterator(nil, nil, 1)
	}
	return q.iterator
}

func (q *Queue) Enqueue(c NodeContent) {
	m := &Node{Data: c}
	i := q.GetIterator()
	if i.head == nil {
		i.head = m
		i.current = i.head
		i.tail = m
		return
	}
	m.prev = i.tail
	i.tail.next = m
	i.tail = m
	i.count++
}

func (q *Queue) Dequeue(strategy DequeueStrategy) NodeContent {
	return strategy.Dequeue(q)
}
