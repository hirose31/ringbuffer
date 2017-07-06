package ringbuffer

import (
	"fmt"
	"reflect"
	"testing"
)

func ck(label string, t *testing.T, rb *RingBuffer, head, tail, len int, buf []interface{}) {
	result := true

	if rb.head != head {
		result = false
	}

	if rb.tail != tail {
		result = false
	}

	if rb.len != len {
		result = false
	}

	if result == false {
		t.Errorf("%s - actual:%v expected:head=%d tail=%d len=%d",
			label,
			rb,
			head,
			tail,
			len,
		)
	}

	if !reflect.DeepEqual(rb.buf, buf) {
		r := reflect.ValueOf(rb.buf)
		fmt.Printf("%s\n", r.Type())
		r = reflect.ValueOf(buf)
		fmt.Printf("%s\n", r.Type())

		t.Errorf("%s - buf actual:%v expected:%v",
			label,
			rb.buf,
			buf,
		)
	}
}

func is(label string, t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("%s - val actual:%s expected:%s",
			label,
			actual,
			expected,
		)
	}
}

func isDeeply(label string, t *testing.T, actual, expected []interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%s - slice actual:%v expected:%v",
			label,
			actual,
			expected,
		)
	}
}

func TestBasic(t *testing.T) {
	rb := NewRingBuffer(5)
	var v interface{}
	var val string
	var vals []interface{}

	ck("init", t, rb, 0, 0, 0, []interface{}{nil, nil, nil, nil, nil})

	rb.Push("1")
	rb.Push("2")
	ck("2 push", t, rb, 0, 2, 2, []interface{}{"1", "2", nil, nil, nil})

	v, _ = rb.Shift()
	val = v.(string)
	is("1 shift", t, val, "1")
	ck("1 shift", t, rb, 1, 2, 1, []interface{}{"1", "2", nil, nil, nil})

	rb.Push("3")
	rb.Push("4")

	vals, _ = rb.Clear()
	isDeeply("clear", t, vals, []interface{}{"2", "3", "4"})
	ck("clear", t, rb, 4, 4, 0, []interface{}{"1", "2", "3", "4", nil})
}

func TestCircular(t *testing.T) {
	rb := NewRingBuffer(5)
	var v interface{}
	var val string
	var vals []interface{}

	rb.Push("1")
	rb.Push("2")
	rb.Push("3")
	rb.Push("4")
	ck("4 push", t, rb, 0, 4, 4, []interface{}{"1", "2", "3", "4", nil})

	rb.Push("5")
	ck("1 push", t, rb, 0, 0, 5, []interface{}{"1", "2", "3", "4", "5"})

	rb.Push("11")
	ck("1 push", t, rb, 1, 1, 5, []interface{}{"11", "2", "3", "4", "5"})

	rb.Push("22")
	ck("1 push", t, rb, 2, 2, 5, []interface{}{"11", "22", "3", "4", "5"})

	v, _ = rb.Shift()
	val = v.(string)
	is("1 shift", t, val, "3")
	ck("1 shift", t, rb, 3, 2, 4, []interface{}{"11", "22", "3", "4", "5"})

	v, _ = rb.Shift()
	val = v.(string)
	is("1 shift", t, val, "4")
	ck("1 shift", t, rb, 4, 2, 3, []interface{}{"11", "22", "3", "4", "5"})

	v, _ = rb.Shift()
	val = v.(string)
	is("1 shift", t, val, "5")
	ck("1 shift", t, rb, 0, 2, 2, []interface{}{"11", "22", "3", "4", "5"})

	v, _ = rb.Shift()
	val = v.(string)
	is("1 shift", t, val, "11")
	ck("1 shift", t, rb, 1, 2, 1, []interface{}{"11", "22", "3", "4", "5"})

	v, _ = rb.Shift()
	val = v.(string)
	is("1 shift", t, val, "22")
	ck("1 shift", t, rb, 2, 2, 0, []interface{}{"11", "22", "3", "4", "5"})

	vals, _ = rb.Clear()
	isDeeply("clear", t, vals, []interface{}{})
	ck("clear", t, rb, 2, 2, 0, []interface{}{"11", "22", "3", "4", "5"})
}

func TestClear(t *testing.T) {
	rb := NewRingBuffer(5)
	var vals []interface{}

	rb.Push("1")
	rb.Push("2")
	rb.Push("3")
	rb.Push("4")
	vals, _ = rb.Clear()
	isDeeply("4 push, clear", t, vals, []interface{}{"1", "2", "3", "4"})
	ck("within size", t, rb, 4, 4, 0, []interface{}{"1", "2", "3", "4", nil})

	rb = NewRingBuffer(5)
	rb.Push("1")
	rb.Push("2")
	rb.Push("3")
	rb.Push("4")
	rb.Push("5")
	vals, _ = rb.Clear()
	isDeeply("5 push, clear", t, vals, []interface{}{"1", "2", "3", "4", "5"})
	ck("equal size", t, rb, 0, 0, 0, []interface{}{"1", "2", "3", "4", "5"})

	rb = NewRingBuffer(5)
	rb.Push("1")
	rb.Push("2")
	rb.Push("3")
	rb.Push("4")
	rb.Push("5")
	rb.Push("11")
	vals, _ = rb.Clear()
	isDeeply("6 push, clear", t, vals, []interface{}{"2", "3", "4", "5", "11"})
	ck("over size", t, rb, 1, 1, 0, []interface{}{"11", "2", "3", "4", "5"})
}
