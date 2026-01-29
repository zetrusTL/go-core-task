package main

import "testing"

func TestAddAndGet(t *testing.T) {
	sm := NewStringIntMap()
	sm.Add("a", 1)

	v, ok := sm.Get("a")
	if !ok || v != 1 {
		t.Fatal("Get after Add failed")
	}
}

func TestExists(t *testing.T) {
	sm := NewStringIntMap()
	sm.Add("x", 5)

	if !sm.Exists("x") {
		t.Fatal("Exists should be true")
	}
	if sm.Exists("y") {
		t.Fatal("Exists should be false")
	}
}

func TestRemove(t *testing.T) {
	sm := NewStringIntMap()
	sm.Add("k", 7)
	sm.Remove("k")

	if sm.Exists("k") {
		t.Fatal("Remove failed")
	}
}

func TestCopyIndependent(t *testing.T) {
	sm := NewStringIntMap()
	sm.Add("a", 1)

	cp := sm.Copy()
	cp["a"] = 999 
	v, _ := sm.Get("a")
	if v == 999 {
		t.Fatal("Copy must be independent")
	}
}
