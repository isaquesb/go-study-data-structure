package main_test

import (
	queue "github.com/isaquesb/go-study-data-structure"
	"testing"
)

func TestLifo_Dequeue(t *testing.T) {
	q := &queue.Queue{}
	q.Enqueue("Alex")
	q.Enqueue("John")
	q.Enqueue("Mary")

	strategy := &queue.Lifo{}

	if current := q.Dequeue(strategy); current != "Mary" {
		t.Errorf("Current got %s, wanted %s", current, "Mary")
	}

	if current := q.Dequeue(strategy); current != "John" {
		t.Errorf("Current got %q, wanted %q", current, "John")
	}

	if current := q.Dequeue(strategy); current != "Alex" {
		t.Errorf("Current got %s, wanted %s", current, "Alex")
	}

	if current := q.Dequeue(strategy); current != nil {
		t.Errorf("Current got %s, wanted nil", current)
	}
}
