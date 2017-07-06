# RingBuffer

simple ring (circular) buffer implementation

<p align="center">
  <!-- a href="https://github.com/hirose31/ringbuffer/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/hirose31/ringbuffer.svg?style=flat-square"></a -->
  <a href="/LICENSE.md"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
  <a href="https://travis-ci.org/hirose31/ringbuffer"><img alt="Travis" src="https://img.shields.io/travis/hirose31/ringbuffer.svg?style=flat-square"></a>
  <a href="https://codecov.io/gh/hirose31/ringbuffer"><img alt="Codecov branch" src="https://img.shields.io/codecov/c/github/hirose31/ringbuffer/master.svg?style=flat-square"></a>
  <a href="https://goreportcard.com/report/github.com/hirose31/ringbuffer"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/hirose31/ringbuffer?style=flat-square"></a>
  <a href="http://godoc.org/github.com/hirose31/ringbuffer"><img alt="Go Doc" src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
</p>

## Installation

```
go get -u github.com/hirose31/ringbuffer
```

## Usage

``` go
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

```

