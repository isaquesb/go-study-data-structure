package main

type Lifo struct {
}

func (l *Lifo) pull(p Piped) NodeContent {
	i := p.GetIterator()
	if i.tail == nil {
		return nil
	}
	m := i.tail
	p.setTail(m.prev)
	i.count--
	return m.Data
}
