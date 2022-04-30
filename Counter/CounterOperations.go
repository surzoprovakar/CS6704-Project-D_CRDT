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
	fmt.Println()
}

func (c *Counter) ToByteArray() []byte {

	id := make([]byte, 64)
	i1 := make([]byte, 64)
	i2 := make([]byte, 64)
	i3 := make([]byte, 64)
	d1 := make([]byte, 64)
	d2 := make([]byte, 64)
	d3 := make([]byte, 64)

	binary.LittleEndian.PutUint64(id, uint64(c.Id()))
	binary.LittleEndian.PutUint64(i1, uint64(c.inc_list[0]))
	binary.LittleEndian.PutUint64(i2, uint64(c.inc_list[1]))
	binary.LittleEndian.PutUint64(i3, uint64(c.inc_list[2]))
	binary.LittleEndian.PutUint64(d1, uint64(c.dec_list[0]))
	binary.LittleEndian.PutUint64(d2, uint64(c.dec_list[1]))
	binary.LittleEndian.PutUint64(d3, uint64(c.dec_list[2]))

	id = append(id, i1...)
	id = append(id, i2...)
	id = append(id, i3...)
	id = append(id, d1...)
	id = append(id, d2...)
	id = append(id, d3...)

	return id
}

func FromByteArray(bytes []byte) *Counter {

	var id = binary.LittleEndian.Uint64(bytes[0 : len(bytes)/7])
	var i1 = binary.LittleEndian.Uint64(bytes[len(bytes)/7 : 2*len(bytes)/7])
	var i2 = binary.LittleEndian.Uint64(bytes[2*len(bytes)/7 : 3*len(bytes)/7])
	var i3 = binary.LittleEndian.Uint64(bytes[3*len(bytes)/7 : 4*len(bytes)/7])
	var d1 = binary.LittleEndian.Uint64(bytes[4*len(bytes)/7 : 5*len(bytes)/7])
	var d2 = binary.LittleEndian.Uint64(bytes[5*len(bytes)/7 : 6*len(bytes)/7])
	var d3 = binary.LittleEndian.Uint64(bytes[6*len(bytes)/7:])

	c := NewCounter(int(id))
	c.inc_list[0] = int(int64(i1))
	c.inc_list[1] = int(int64(i2))
	c.inc_list[2] = int(int64(i3))
	c.dec_list[0] = int(int64(d1))
	c.dec_list[1] = int(int64(d2))
	c.dec_list[2] = int(int64(d3))
	return c
}

func (c *Counter) Merge(o *Counter) {
	mu.Lock()
	fmt.Println("Starting to merge......")
	fmt.Println("Counter Itself")
	c.Print()
	fmt.Println("Requester Counter")
	o.Print()

	c.inc_list[o.id-1] = o.inc_list[o.id-1]
	c.dec_list[o.id-1] = o.dec_list[o.id-1]

	fmt.Println("Merged")
	c.Print()
	mu.Unlock()
}
