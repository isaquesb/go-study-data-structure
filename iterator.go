package main

type Iterator struct {
	head    *Message
	tail    *Message
	current *Message
	count   int
}

func NewIterator(head *Message, tail *Message, count int) *Iterator {
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

func (i *Iterator) Current() MessageContent {
	return i.current.Data
}

func (i *Iterator) Next() {
	if i.current != nil {
		i.current = i.current.next
	}
}

func (i *Iterator) Valid() bool {
	return i.current != nil
}
