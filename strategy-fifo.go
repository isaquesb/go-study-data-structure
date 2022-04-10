package main

type Fifo struct {
}

func (f *Fifo) pull(p Piped) NodeContent {
	i := p.GetIterator()
	if i.head == nil {
		return nil
	}
	m := i.head
	p.setHead(m.next)
	i.count--
	return m.Data
}
