package main

type LinkedList struct {
	pipe
}

func (l *LinkedList) AddFirst(c NodeContent) {
	l.addFirst(c)
}

func (l *LinkedList) AddLast(c NodeContent) {
	l.addLast(c)
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
