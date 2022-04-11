package main_test

import (
	ds "github.com/isaquesb/go-study-data-structure"
	"testing"
)

func TestHashMap_Has(t *testing.T) {
	k := "user"
	h := ds.NewHashMap(ds.HashConfig{})

	if h.Has(k) {
		t.Errorf("has not key failed")
	}

	h.Put(k, "John")

	if !h.Has(k) {
		t.Errorf("inputted key \"%s\" not exists", k)
	}
}

func TestHashMap_Delete(t *testing.T) {
	k := "contact"
	n := "Mary"
	h := ds.NewHashMap(ds.HashConfig{Size: 1})
	h.Put(k, n)

	if !h.Has(k) {
		t.Errorf("inputted key \"%s\" not exists", k)
	}

	deleted, name := h.Delete(k)

	if !deleted {
		t.Errorf("error on delete key \"%s\"", k)
	}

	if name != n {
		t.Errorf("name got %s, expected \"%s\"", name, n)
	}

	deleted, name = h.Delete(k)

	if deleted {
		t.Errorf("error on delete deleted key \"%s\"", k)
	}

	if nil != name {
		t.Errorf("name got %s, expected nil", name)
	}
}

func TestHashMap_Search(t *testing.T) {
	u := struct {
		name string
		age  int
	}{"John", 38}
	h := ds.NewHashMap(ds.HashConfig{Size: 2})

	has, name := h.Search("name")

	if has {
		t.Error("has unexpected name key")
	}

	if nil != name {
		t.Errorf("got unexpected name: %s", name)
	}

	h.Put("name", u.name)
	h.Put("age", u.age)

	if _, name := h.Search("name"); name != u.name {
		t.Errorf("name got %s, expected \"%s\"", name, u.name)
	}

	if _, age := h.Search("age"); age != u.age {
		t.Errorf("age got %q, expected \"%d\"", age, u.age)
	}
}
