package main_test

import (
	ds "github.com/isaquesb/go-study-data-structure"
	"testing"
)

func TestStack_Push(t *testing.T) {
	q := &ds.Stack{}
	q.Push("Alex")
	q.Push("John")
	q.Push("Mary")

	i := q.GetIterator()
	count := i.Count()
	want := 3
	if count != want {
		t.Errorf("got %q, expected %q", count, want)
	}
}

func TestStack_Pop(t *testing.T) {
	q := &ds.Stack{}
	q.Push("Alex")
	q.Push("John")
	q.Push("Mary")

	i := q.GetIterator()

	for _, v := range []struct {
		n string
		c int
	}{
		{"Mary", 3},
		{"John", 2},
		{"Alex", 1},
	} {
		if v.c != i.Count() {
			t.Errorf("got %d, expected %d", i.Count(), v.c)
		}
		if name := q.Pop(); name != v.n {
			t.Errorf("First got %q, expected %q", name, v.n)
		}
	}
	if name := q.Pop(); nil != name {
		t.Error("Queue problem")
	}
}
