# fifo

Fifo queue go package implemented as growable circular buffer.
This is the equivalent of the deque data structure found in the C++ STL.

## General properties 

Elements may be added to the front or the back of the queue, and removed 
from both ends.

The queue is implemented using an array of interfaces which should put less stress 
to the garbage collector than with list containers. Access is not goroutine safe 
(no Mutex locks) and should thus be more efficient than channels. 

When an element is added to a full fifo queue, the buffer size is doubled and the
elements are copied into the new buffer. This should be ok for most use cases but
could be a problem when the queue contains a huge amount of elements because the
alement addition operation would block until growing the buffer is completed. 
It is possible to avoid this problem by spreading the grow operation on multiple
element additions, but it would make the code more complex and less efficient for
simple usage cases which we currently assume to be more common. 

Use the following command to install the package:

    go get "github.com/chmike/fifo"
	
## Fifo API

To use the package add the following import instruction 

    import "github.com/chmike/fifo"
	
Values are stored as `interface{}` values. When obtaining a value from a fifo 
element access method, dynamic typing handling is required cast the value to the
expected type. A typical example is as follow:

	f := fifo.New(0)
    var x MyType = ...
    f.PushBack(x)
	. . .
	v, ok := f.Back().(MyType)
	if ok {
		// The back value is of type MyType and v holds a copy of it, unless
		// MyType is an interface in which case v holds a reference to the object.
	} else {
		// The back value in the fifo is not of type MyType 
	}
		

### func New(capacity int) *Fifo

`New` creates a new fifo queue with the specified capacity.
A value smaller than the default inital capacity is ignored.  

    f := fifo.New(0)
	
### func (*Fifo) Init(capacity int) *Fifo

`Init` initializes the fifo queue with the specified capacity, or clears the 
content of the fifo. A value smaller than the default inital capacity is ignored.  
On return the fifo capacity is at least the specified capacity. 

    f.Init(100)

### func (*Fifo) Len() int

`Len` returns the number of elements in the fifo queue.

    n := f.Len()

### func (*Fifo) PushFront(interface{}) interface{}

`PushFront` insert the argument at the front of the fifo queue and returns it. 
Use dynamic casting to validate the type of the returned value.

    v, ok := f.PushFront(5).(int)

### func (*Fifo) PushBack(interface{}) interface{}

`PushBack` insert the argument at the back of the fifo queue and returns it. 
Use dynamic casting to validate the type of the returned value.

    v, ok := f.PushBack(42).(int)

### func (*Fifo) Front() interface{}

`Front` returns the element at the front of the fifo queue or `nil` if empty. 
Use dynamic casting to validate the type of the returned value.

    v, ok := f.Front().(int)

### func (*Fifo) Back() interface{}

`Back` returns the element at the back of the fifo queue or `nil` if empty. 
Use dynamic casting to validate the type of the returned value.

    v, ok := f.Back().(int)

### func (*Fifo) PopFront() interface{}

`PopFront` removes the element at the front of the fifo queue and returns it. 
Use dynamic casting to validate the type of the returned value.

	f.PopFront() // ignoring the return value
    v, ok := f.PopFront().(int)

### func (*Fifo) PopBack() interface{}

`PopBack` removes the element at the back of the fifo queue and returns it.  
Use dynamic casting to validate the type of the returned value.

    f.PopBack() // ignoring the return value
	v, ok := f.PopBack().(int)



