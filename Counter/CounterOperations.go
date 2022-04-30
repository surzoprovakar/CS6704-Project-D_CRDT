package main

import (
	"encoding/binary"
	"fmt"
	"sync"
)

var mu sync.Mutex

type Counter struct {
	id       int
	value    int
	inc_list []int
	dec_list []int
}

func NewCounter(id int) *Counter {
	return &Counter{id: id, value: 0, inc_list: make([]int, 3), dec_list: make([]int, 3)}
}

func (c *Counter) Id() int {
	return c.id
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) IncList() []int {
	return c.inc_list
}

func (c *Counter) DecList() []int {
	return c.dec_list
}

func (c *Counter) Inc() {
	mu.Lock()
	c.value++
	mu.Unlock()
}

func (c *Counter) Dec() {
	mu.Lock()
	c.value--
	mu.Unlock()
}

func (c *Counter) Print() {
	fmt.Println("Counter: ", c.id)
	fmt.Println(c.inc_list)
	fmt.Println(c.dec_list)
}

func (c *Counter) ToByteArray() []byte {

	a1 := make([]byte, 64)
	a2 := make([]byte, 64)

	binary.LittleEndian.PutUint64(a1, uint64(c.Id()))
	binary.LittleEndian.PutUint64(a2, uint64(c.Value()))
	return append(a1, a2...)
}

func FromByteArray(bytes []byte) *Counter {

	var r1 = binary.LittleEndian.Uint64(bytes[0:(len(bytes) / 2)])
	var r2 = binary.LittleEndian.Uint64(bytes[len(bytes)/2:])
	id := int64(r1)
	c := NewCounter(int(id))
	c.value = int(int64(r2))
	return c

}
