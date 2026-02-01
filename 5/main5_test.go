package main

import (
	"reflect"
	"testing"
)

func TestDiffHasIntersection(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	ok, res := Diff(a, b)

	if !ok {
		t.Fatal("expected true, got false")
	}

	expected := []int{64, 3}

	if !reflect.DeepEqual(res, expected) {
		t.Fatal("unexpected result:", res)
	}
}

func TestDiffNoIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	ok, res := Diff(a, b)

	if ok {
		t.Fatal("expected false, got true")
	}

	if len(res) != 0 {
		t.Fatal("expected empty slice")
	}
}
