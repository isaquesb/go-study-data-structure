package main

type Fifo struct {
	StrategyAbstract
}

func (f *Fifo) Dequeue(q *Queue) NodeContent {
	i := q.GetIterator()
	if i.head == nil {
		return nil
	}
	m := i.head
	f.setHead(m.next, q)
	i.count--
	return m.Data
}
