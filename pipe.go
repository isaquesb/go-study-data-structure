package main

type Piped interface {
	GetIterator() *Iterator
	setHead(m *Node)
	setTail(m *Node)
}

type pipe struct {
	iterator *Iterator
}

func (p *pipe) GetIterator() *Iterator {
	if p.iterator == nil {
		p.iterator = NewIterator(nil, nil, 1)
	}
	return p.iterator
}

func (p *pipe) add(c NodeContent) {
	m := &Node{Data: c}
	i := p.GetIterator()
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

func (p *pipe) setHead(m *Node) {
	iterator := p.GetIterator()
	iterator.head = m
	if iterator.head != nil {
		iterator.head.prev = nil
	} else {
		p.setTail(nil)
	}
}

func (p *pipe) setTail(m *Node) {
	iterator := p.GetIterator()
	iterator.tail = m
	if iterator.tail != nil {
		iterator.tail.next = nil
	}
}