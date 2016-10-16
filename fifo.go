// Package fifo implements a fifo queue using a circular buffer.
// It support efficient element access, queuing and dequeuing at both ends,
// but no insertion.
// The internal buffer will grow as needed.
package fifo

// Fifo represents a growable circular buffer.
type Fifo struct {
	buf                   []interface{}
	front, back, len, cap int
}

const minCapacity int = 16

// Init initializes or clears fifo f.
// The fifo queue is guarateed to have at least the specified capacity
func (f *Fifo) Init(capacity int) *Fifo {
	if capacity > f.cap {
		f.cap = capacity
	}
	if f.cap < minCapacity {
		f.cap = minCapacity
	}
	f.buf = make([]interface{}, f.cap)
	f.front = 0
	f.back = 0
	f.len = 0
	return f
}

// New returns an initialized fifo with the specified initial capacity.
func New(capacity int) *Fifo {
	return new(Fifo).Init(capacity)
}

// Len returns the number of elements in the fifo queue.
// The complexity is O(1)
func (f *Fifo) Len() int {
	return f.len
}

// Front returns the front element in the fifo queue or nil if the fifo is empty.
func (f *Fifo) Front() interface{} {
	if f.len == 0 {
		return nil
	}
	return f.buf[f.front]
}

// Back returns the back element in the fifo queue or nil if the fifo is empty.
func (f *Fifo) Back() interface{} {
	if f.len == 0 {
		return nil
	}
	if f.back > 0 {
		return f.buf[f.back-1]
	}
	return f.buf[f.cap-1]
}

// PopFront returns the front element in the fifo queue after removing it from the fifo.
// It returns nil if the fifo queue is empty.
func (f *Fifo) PopFront() interface{} {
	if f.len == 0 {
		return nil
	}
	front := f.buf[f.front]
	f.buf[f.front] = nil // for garbage collection
	f.front++
	if f.front == f.cap {
		f.front = 0
	}
	f.len--
	return front
}

// PopFront returns the back element in the fifo queue after removing it from the fifo.
// It returns nil if the fifo queue is empty.
func (f *Fifo) PopBack() interface{} {
	if f.len == 0 {
		return nil
	}
	if f.back == 0 {
		f.back = f.cap - 1
	} else {
		f.back--
	}
	back := f.buf[f.back]
	f.buf[f.back] = nil
	f.len--
	return back
}

// growBuffer doubles the buffer capacity when adding an element would overflow it.
func (f *Fifo) growBuffer() {
	// Assume buffer is full : f.len == l.cap && f.back == f.front
	newBuf := make([]interface{}, f.cap*2)
	j := 0
	for i := f.front; i < f.cap; i++ {
		newBuf[j] = f.buf[i]
		j++
	}
	for i := 0; i < f.back; i++ {
		newBuf[j] = f.buf[i]
		j++
	}
	f.buf = newBuf
	f.cap = len(newBuf)
	f.front = 0
	f.back = j
}

// PushFront insert e at the front of the fifo queue and returns e.
func (f *Fifo) PushFront(e interface{}) interface{} {
	if f.len == f.cap {
		f.growBuffer()
	}
	if f.front == 0 {
		f.front = f.cap - 1
	} else {
		f.front--
	}
	f.buf[f.front] = e
	f.len++
	return e
}

// PushBack insert e at the back of the fifo queue and returns e.
func (f *Fifo) PushBack(e interface{}) interface{} {
	if f.len == f.cap {
		f.growBuffer()
	}
	f.buf[f.back] = e
	f.back++
	if f.back == f.cap {
		f.back = 0
	}
	f.len++
	return e
}
