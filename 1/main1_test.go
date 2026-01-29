package main

import "testing"

func TestConcatToString(t *testing.T) {
	res := ConcatToString(1, "a", true)
	if res != "1atrue" {
		t.Fatal("unexpected result:", res)
	}
}

func TestStringToRunes(t *testing.T) {
	r := StringToRunes("go")
	if len(r) != 2 {
		t.Fatal("wrong rune length")
	}
}

func TestHashRunesWithSalt(t *testing.T) {
	r := []rune("test")
	h1 := HashRunesWithSalt(r, "go-2024")
	h2 := HashRunesWithSalt(r, "go-2024")

	if h1 != h2 {
		t.Fatal("hash must be deterministic")
	}
}
