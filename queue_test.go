package main_test

import (
	ds "github.com/isaquesb/go-study-data-structure"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	q := &ds.Queue{}
	q.Enqueue("Alex")
	q.Enqueue("John")
	q.Enqueue("Mary")

	i := q.GetIterator()
	count := i.Count()
	want := 3
	if count != want {
		t.Errorf("got %q, expected %q", count, want)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := &ds.Queue{}
	q.Enqueue("Alex")
	q.Enqueue("John")
	q.Enqueue("Mary")

	i := q.GetIterator()

	for _, v := range []struct {
		n string
		c int
	}{
		{"Alex", 3},
		{"John", 2},
		{"Mary", 1},
	} {
		if v.c != i.Count() {
			t.Errorf("got %d, expected %d", i.Count(), v.c)
		}
		if name := q.Dequeue(); name != v.n {
			t.Errorf("First got %q, expected %q", name, v.n)
		}
	}
	if name := q.Dequeue(); nil != name {
		t.Error("Queue problem")
	}
}
