package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}

	result := sliceExample(input)

	if !reflect.DeepEqual(result, expected) {
		t.Fatal("sliceExample failed:", result)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	result := addElements(input, 10)

	expected := []int{1, 2, 3, 10}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal("addElements failed:", result)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{5, 6, 7}
	copied := copySlice(input)

	input[0] = 999

	if copied[0] == 999 {
		t.Fatal("copySlice is not independent")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4}

	result := removeElement(input, 1)
	expected := []int{1, 3, 4}

	if !reflect.DeepEqual(result, expected) {
		t.Fatal("removeElement failed:", result)
	}
}
