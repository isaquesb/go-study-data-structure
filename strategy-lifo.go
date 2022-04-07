package main

type Lifo struct {
	StrategyAbstract
}

func (l *Lifo) Dequeue(q *Queue) NodeContent {
	i := q.GetIterator()
	if i.tail == nil {
		return nil
	}
	m := i.tail
	l.setTail(m.prev, q)
	i.count--
	return m.Data
}
