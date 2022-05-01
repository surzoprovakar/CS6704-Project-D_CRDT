package main

import (
	"encoding/binary"
	"fmt"
	"sort"
	"sync"
)

var mu sync.Mutex

type Set struct {
	id     int
	values map[int]struct{}
}

func NewSet(id int) *Set {
	return &Set{id: id, values: make(map[int]struct{})}
}

func (s *Set) Id() int {
	return s.id
}

func (s *Set) Values() map[int]struct{} {
	return s.values
}

func (s *Set) Contain(value int) bool {
	_, c := s.values[value]
	return c
}

func (s *Set) Add(value int) {
	mu.Lock()
	if !s.Contain(value) {
		s.values[value] = struct{}{}
	}
	mu.Unlock()
}

func (s *Set) Remove(value int) {
	mu.Lock()
	if s.Contain(value) {
		delete(s.values, value)
	}
	mu.Unlock()
}

func (s *Set) ToByteArray() []byte {

	a1 := make([]byte, 64)

	binary.LittleEndian.PutUint64(a1, uint64(s.Id()))

	values := get_keys_from_map(s.values)
	for i := range values {
		tmp := make([]byte, 64)
		binary.LittleEndian.PutUint64(tmp, uint64(values[i]))
		a1 = append(a1, tmp...)
	}
	return a1
}

func FromByteArray(bytes []byte) *Set {

	var length = len(bytes)
	var div = length / 64

	var r1 = binary.LittleEndian.Uint64(bytes[0:(len(bytes) / div)])

	i := 1
	var tmp_list []int
	for i < div {
		var r2 = binary.LittleEndian.Uint64(bytes[i*len(bytes)/div : (i+1)*len(bytes)/div])
		tmp_list = append(tmp_list, int(r2))
		i++
	}

	id := int64(r1)
	s := NewSet(int(id))

	for j := range tmp_list {
		s.values[tmp_list[j]] = struct{}{}
	}

	return s

}

func (s *Set) Print() {
	fmt.Print("Set:", s.id, " ")
	values := get_keys_from_map(s.values)
	fmt.Println(values)
}

func get_keys_from_map(mapp map[int]struct{}) []int {
	keys := make([]int, 0, len(mapp))

	for key := range mapp {
		keys = append(keys, key)
	}
	sort.Ints(keys[:])
	return keys
}

func (s *Set) Merge(o *Set) {

	mu.Lock()
	fmt.Println("Starting to merge:")
	s.Print()
	o.Print()
	m := make(map[int]bool)

	for item := range s.values {
		m[item] = true
	}

	for item := range o.values {
		if _, ok := m[item]; !ok {
			s.values[item] = struct{}{}
		}
	}
	fmt.Print("Merged ")
	s.Print()
	mu.Unlock()
}
