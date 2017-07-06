// +build ignore

package main

import (
	"fmt"

	"github.com/hirose31/ringbuffer"
)

func main() {
	rb := ringbuffer.NewRingBuffer(3)

	// push & shift
	rb.Push("foo")
	rb.Push("bar")

	val, err := rb.Shift()
	if err != nil {
		panic(err)
	}
	fmt.Printf("got: %s\n", val.(string))
	// => got: foo

	// clear and fetch all elements
	rb.Push("baz")
	vals, err := rb.Clear()
	if err != nil {
		panic(err)
	}
	for i, v := range vals {
		fmt.Printf("[%d]%s\n", i, v)
	}
	// => [0]bar
	// => [1]baz

	// circular
	rb.Push("one")
	rb.Push("two")
	rb.Push("three")
	rb.Push("four")
	vals, err = rb.Clear()
	if err != nil {
		panic(err)
	}
	for i, v := range vals {
		fmt.Printf("[%d]%s\n", i, v)
	}
	// "one" is overwritten with "four"
	// => [0]two
	// => [1]three
	// => [2]four
}
