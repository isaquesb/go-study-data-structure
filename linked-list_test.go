package main_test

import (
	ds "github.com/isaquesb/go-study-data-structure"
	"testing"
)

func TestLinkedList_AddLast(t *testing.T) {
	list := &ds.LinkedList{}
	list.AddLast("Alex")
	list.AddLast("John")
	list.AddLast("Mary")

	i := list.GetIterator()

	for j, v := range []string{"Alex", "John", "Mary"} {
		if c := i.Current(); c != v {
			t.Errorf("scenery %d: got %s, expected %s", j, c, v)
		}
		i.Next()
	}
}

func TestLinkedList_AddFirst(t *testing.T) {
	list := &ds.LinkedList{}

	if first := list.GetFirst(); first != nil {
		t.Errorf("first got %s, expected nil", first)
		return
	}

	if last := list.GetLast(); last != nil {
		t.Errorf("last got %s, expected nil", last)
		return
	}

	list.AddFirst("Alex")
	list.AddFirst("John")
	list.AddFirst("Mary")

	if first := list.GetFirst(); first != "Mary" {
		t.Errorf("fist got %s, expected %s", first, "Mary")
		return
	}

	if last := list.GetLast(); last != "Alex" {
		t.Errorf("last got %s, expected %s", last, "Alex")
		return
	}

	i := list.GetIterator()

	for j, v := range []string{"Mary", "John", "Alex"} {
		if c := i.Current(); c != v {
			t.Errorf("scenery %d: got %s, expected %s", j, c, v)
		}
		i.Next()
	}
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	list := &ds.LinkedList{}
	list.AddFirst("Alex")
	list.AddLast("John")
	list.AddFirst("Mary")

	list.RemoveFirst()

	i := list.GetIterator()

	for j, v := range []string{"Alex", "John"} {
		if c := i.Current(); c != v {
			t.Errorf("scenery %d: got %s, expected %s", j, c, v)
		}
		i.Next()
	}
}

func TestLinkedList_RemoveLast(t *testing.T) {
	list := &ds.LinkedList{}
	list.AddFirst("Alex")
	list.AddLast("John")
	list.AddFirst("Mary")

	if last := list.GetLast(); last != "John" {
		t.Errorf("last got %s, expected %s", last, "John")
		return
	}

	list.RemoveLast()

	i := list.GetIterator()
	i.Rewind()

	for j, v := range []string{"Mary", "Alex"} {
		if c := i.Current(); c != v {
			t.Errorf("scenery %d: got %s, expected %s", j, c, v)
		}
		i.Next()
	}

	if i.Valid() {
		t.Errorf("Iterator problem")
	}
}
