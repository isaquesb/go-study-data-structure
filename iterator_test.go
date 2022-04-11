package main_test

import (
	ds "github.com/isaquesb/go-study-data-structure"
	"testing"
)

func TestIterator_Navigation(t *testing.T) {
	q := &ds.Queue{}
	i := q.GetIterator()

	if c := i.Current(); nil != c {
		t.Errorf("got %s, expected nil", c)
		return
	}

	q.Enqueue("Alex")
	q.Enqueue("John")
	q.Enqueue("Mary")

	if c := i.Current(); c != "Alex" {
		t.Errorf("got %s, expected %s", c, "Alex")
		return
	}
	i.Rewind()

	for j, v := range []string{"Alex", "John", "Mary"} {
		if c := i.Current(); c != v {
			t.Errorf("scenery %d: got %s, expected %s", j, c, v)
		}
		i.Next()
	}

	if i.Valid() {
		t.Errorf("Iterator problem")
	}
}
