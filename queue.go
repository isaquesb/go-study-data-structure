package main

type MessageContent = interface{}

type DequeueStrategy interface {
	Dequeue(q *Queue) MessageContent
}

type Message struct {
	Data MessageContent
	next *Message
	prev *Message
}

type Queue struct {
	iterator *Iterator
}

type Strategy struct {
}

type Fifo struct {
	Strategy
}

type Lifo struct {
	Strategy
}

func (q *Queue) GetIterator() *Iterator {
	if q.iterator == nil {
		q.iterator = NewIterator(nil, nil, 1)
	}
	return q.iterator
}

func (q *Queue) Enqueue(c MessageContent) {
	m := &Message{Data: c}
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

func (q *Queue) Dequeue(strategy DequeueStrategy) MessageContent {
	return strategy.Dequeue(q)
}

func (f *Strategy) setHead(m *Message, q *Queue) {
	q.iterator.head = m
	if q.iterator.head != nil {
		q.iterator.head.prev = nil
	} else {
		f.setTail(nil, q)
	}
}

func (f *Strategy) setTail(m *Message, q *Queue) {
	q.iterator.tail = m
	if q.iterator.tail != nil {
		q.iterator.tail.next = nil
	}
}

func (f *Fifo) Dequeue(q *Queue) MessageContent {
	i := q.GetIterator()
	if i.head == nil {
		return nil
	}
	m := i.head
	f.setHead(m.next, q)
	i.count--
	return m.Data
}

func (l *Lifo) Dequeue(q *Queue) MessageContent {
	i := q.GetIterator()
	if i.tail == nil {
		return nil
	}
	m := i.tail
	l.setTail(m.prev, q)
	i.count--
	return m.Data
}
