package main_test

import (
	queue "github.com/isaquesb/go-study-data-structure"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	q := &queue.Queue{}
	q.Enqueue("Alex")
	q.Enqueue("John")
	q.Enqueue("Mary")

	i := q.GetIterator()
	count := i.Count()
	want := 3
	if count != want {
		t.Errorf("got %q, wanted %q", count, want)
	}
}

func TestQueue_List(t *testing.T) {
	q := &queue.Queue{}
	q.Enqueue("Alex")
	q.Enqueue("John")
	q.Enqueue("Mary")

	i := q.GetIterator()

	if c := i.Current(); c != "Alex" {
		t.Errorf("got %s, wanted %s", c, "Alex")
	}
	i.Next()
	if c := i.Current(); c != "John" {
		t.Errorf("got %s, wanted %s", c, "John")
	}
	i.Next()

	if c := i.Current(); c != "Mary" {
		t.Errorf("got %s, wanted %s", c, "Mary")
	}

	i.Rewind()

	if c := i.Current(); c != "Alex" {
		t.Errorf("got %s, wanted %s", c, "Alex")
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := &queue.Queue{}
	q.Enqueue("Alex")
	q.Enqueue("John")
	q.Enqueue("Mary")

	i := q.GetIterator()

	lifo := &queue.Lifo{}

	first := q.Dequeue(&queue.Fifo{})
	count := i.Count()

	if count != 2 {
		t.Errorf("got %d, wanted %d", count, 2)
	}

	last := q.Dequeue(lifo)

	if first != "Alex" {
		t.Errorf("First got %q, wanted %q", first, "Alex")
	}

	if last != "Mary" {
		t.Errorf("Last got %q, wanted %q", last, "Mary")
	}

	if one := q.Dequeue(lifo); one != "John" {
		t.Errorf("Last got %q, wanted %q", one, "John")
	}

	if i.Valid() {
		t.Errorf("Iterator problem")
	}
}
