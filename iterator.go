package main

type Iterator struct {
	head    *Node
	tail    *Node
	current *Node
	count   int
}

func NewIterator(head *Node, tail *Node, count int) *Iterator {
	return &Iterator{
		head:    head,
		tail:    tail,
		current: head,
		count:   count,
	}
}

func (i *Iterator) Count() int {
	return i.count
}

func (i *Iterator) Rewind() {
	i.current = i.head
}

func (i *Iterator) Current() NodeContent {
	if nil != i.current {
		return i.current.Data
	}
	return nil
}

func (i *Iterator) Next() {
	if i.current != nil {
		i.current = i.current.next
	}
}

func (i *Iterator) Valid() bool {
	return i.current != nil && i.head != nil && i.tail != nil
}
