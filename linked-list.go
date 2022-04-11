package main

type LinkedList struct {
	pipe
}

func (l *LinkedList) AddFirst(c NodeContent) {
	m := &Node{Data: c}
	i := l.GetIterator()
	m.next = i.head
	currentIsHead := i.current == i.head
	if nil != i.head {
		i.head.prev = m
	} else {
		i.tail = m
	}
	i.head = m
	i.count++
	if currentIsHead {
		i.current = i.head
	}
}

func (l *LinkedList) AddLast(c NodeContent) {
	l.add(c)
}

func (l *LinkedList) RemoveFirst() {
	strategy := Fifo{}
	strategy.pull(l)
}

func (l *LinkedList) RemoveLast() {
	strategy := Lifo{}
	strategy.pull(l)
}

func (l *LinkedList) GetFirst() NodeContent {
	iterator := l.GetIterator()
	if nil != iterator.head {
		return iterator.head.Data
	}
	return nil
}

func (l *LinkedList) GetLast() NodeContent {
	iterator := l.GetIterator()
	if nil != iterator.tail {
		return iterator.tail.Data
	}
	return nil
}
