// Package ringbuffer is simple ring (circular) buffer implementation.
package ringbuffer

import "fmt"

// RingBuffer type represents ring buffer.
type RingBuffer struct {
	head int
	tail int
	len  int
	size int
	buf  []interface{}
}

// NewRingBuffer returns a new RingBuffer.
func NewRingBuffer(size int) *RingBuffer {
	rb := new(RingBuffer)
	rb.buf = make([]interface{}, size)
	rb.size = size
	return rb
}

// Push pushes the element val to the tail of the ring buffer.
func (rb *RingBuffer) Push(val interface{}) error {
	rb.buf[rb.tail] = val
	rb.len++
	rb.tail = (rb.tail + 1) % rb.size
	if rb.len > rb.size {
		rb.len = rb.size
		rb.head = rb.tail
	}

	return nil
}

// Shift removes an element from head of the ring buffer and returns it.
func (rb *RingBuffer) Shift() (interface{}, error) {
	if rb.len <= 0 {
		return "", fmt.Errorf("%s", "no buffer")
	}

	val := rb.buf[rb.head]
	rb.head = (rb.head + 1) % rb.size
	rb.len--
	return val, nil
}

// Fetch returns all elements
func (rb *RingBuffer) Fetch() ([]interface{}, error) {
	val := make([]interface{}, rb.len)

	len := rb.len
	for i := 0; len > 0; len-- {
		val[i] = rb.buf[(rb.head+i)%rb.size]
		i++
	}

	return val, nil
}

// Clear removes all elements and returns those elements.
func (rb *RingBuffer) Clear() ([]interface{}, error) {
	val := make([]interface{}, rb.len)

	for i := 0; rb.len > 0; rb.len-- {
		val[i] = rb.buf[rb.head]
		i++
		rb.head = (rb.head + 1) % rb.size
	}

	rb.len = 0
	rb.tail = rb.head

	return val, nil
}
