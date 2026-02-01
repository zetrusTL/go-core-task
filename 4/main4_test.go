package main

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	s1 := []string{"apple", "banana", "cherry", "date"}
	s2 := []string{"banana", "date", "fig"}

	got := Diff(s1, s2)
	want := []string{"apple", "cherry"}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("Diff failed:", got)
	}
}