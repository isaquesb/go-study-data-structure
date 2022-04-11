package main_test

import (
	ds "github.com/isaquesb/go-study-data-structure"
	"testing"
)

func TestIterator_Navigation(t *testing.T) {
	q := &ds.Queue{}
	q.Enqueue("Alex")
	q.Enqueue("John")
	q.Enqueue("Mary")

	i := q.GetIterator()

	if c := i.Current(); c != "Alex" {
		t.Errorf("got %s, expected %s", c, "Alex")
	}
	i.Rewind()

	for j, v := range []string{"Alex", "John", "Mary"} {
		if c := i.Current(); c != v {
			t.Errorf("Scenery %d: got %s, expected %s", j, c, v)
		}
		i.Next()
	}

	if i.Valid() {
		t.Errorf("Iterator problem")
	}
}
