package main

import "fmt"

type StringIntMap struct {
	m map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{m: make(map[string]int)}
}

func (s *StringIntMap) Add(key string, value int) {
	s.m[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.m, key)
}

func (s *StringIntMap) Copy() map[string]int {
	cp := make(map[string]int, len(s.m))
	for k, v := range s.m {
		cp[k] = v
	}
	return cp
}

func (s *StringIntMap) Exists(key string) bool {
	_, ok := s.m[key]
	return ok
}

func (s *StringIntMap) Get(key string) (int, bool) {
	v, ok := s.m[key]
	return v, ok
}

func main() {
	sm := NewStringIntMap()

	sm.Add("apple", 10)
	sm.Add("banana", 20)

	fmt.Println("Exists apple:", sm.Exists("apple"))

	v, ok := sm.Get("banana")
	fmt.Println("Get banana:", v, ok)

	cp := sm.Copy()
	fmt.Println("Copy:", cp)

	sm.Remove("apple")
	fmt.Println("After remove apple, Exists apple:", sm.Exists("apple"))
}
