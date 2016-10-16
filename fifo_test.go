package fifo

import (
	"testing"
)

func testEmptyFifo(t *testing.T, f *Fifo) {
	if f.Len() != 0 {
		t.Fatal("Len initialized value is invalid")
	}
	if f.Front() != nil {
		t.Fatal("Front initialized value is invalid")
	}
	if f.Len() != 0 {
		t.Fatal("Len value corrupted by Front() on empty fifo")
	}
	if f.Back() != nil {
		t.Fatal("Back initialized value is invalid")
	}
	if f.Len() != 0 {
		t.Fatal("Len value corrupted by Back() on empty fifo")
	}
	if f.PopFront() != nil {
		t.Fatal("PopFront on empty fifo returned an invalid value")
	}
	if f.Len() != 0 {
		t.Fatal("Len value corrupted by PopFront() on empty fifo")
	}
	if f.PopBack() != nil {
		t.Fatal("PopBack on empty fifo returned an invalid value")
	}
	if f.Len() != 0 {
		t.Fatal("Len value corrupted by PopBack() on empty fifo")
	}
}

func TestEmptyFifo(t *testing.T) {
	testEmptyFifo(t, New(0))
}

func TestNoneEmptyFifo(t *testing.T) {
	const n int = 100000
	f := New(0)
	i := 0
	j := 0
	for i < n {
		for k := 0; k < 9 && i < n; k++ {
			v, ok := f.PushBack(i).(int)
			if !ok {
				t.Fatal("PushBack retuned an invalid value, expected an int")
			} else if v != i {
				t.Fatalf("PushBack retuned %d instead of %d", v, i)
			}
			v, ok = f.Back().(int)
			if !ok {
				t.Fatal("Back retuned an invalid value, expected an int")
			} else if v != i {
				t.Fatalf("Back retuned %d instead of %d", v, i)
			}
			i++
		}
		for k := 0; k < 6; k++ {
			v, ok := f.Front().(int)
			if !ok {
				t.Fatal("Front retuned an invalid value, expected an int")
			} else if v != j {
				t.Fatalf("Front retuned %d instead of %d", v, j)
			}
			v, ok = f.PopFront().(int)
			if !ok {
				t.Fatal("PopFront retuned an invalid value, expected an int")
			} else if v != j {
				t.Fatalf("PopFront retuned %d instead of %d", v, j)
			}
			j++
		}
		if f.Len() != (i+1)-(j+1) {
			t.Fatal("Len of none empty fifo is invalid")
		}
	}
	for j < n {
		v, ok := f.Front().(int)
		if !ok {
			t.Fatal("Front retuned an invalid value, expected an int")
		} else if v != j {
			t.Fatalf("Front retuned %d instead of %d", v, j)
		}
		v, ok = f.PopFront().(int)
		if !ok {
			t.Fatal("PopFront retuned an invalid value, expected an int")
		} else if v != j {
			t.Fatalf("PopFront retuned %d instead of %d", v, j)
		}
		j++
	}
	testEmptyFifo(t, f)

	i = 0
	j = 0
	for i < n {
		for k := 0; k < 9 && i < n; k++ {
			v, ok := f.PushFront(i).(int)
			if !ok {
				t.Fatal("PushFront retuned an invalid value, expected an int")
			} else if v != i {
				t.Fatalf("PushFront retuned %d instead of %d", v, i)
			}
			v, ok = f.Front().(int)
			if !ok {
				t.Fatal("Front retuned an invalid value, expected an int")
			} else if v != i {
				t.Fatalf("Front retuned %d instead of %d", v, i)
			}
			i++
		}
		for k := 0; k < 5; k++ {
			v, ok := f.Back().(int)
			if !ok {
				t.Fatal("Back retuned an invalid value, expected an int")
			} else if v != j {
				t.Fatalf("Back retuned %d instead of %d", v, j)
			}
			v, ok = f.PopBack().(int)
			if !ok {
				t.Fatal("PopBack retuned an invalid value, expected an int")
			} else if v != j {
				t.Fatalf("PopBack retuned %d instead of %d", v, j)
			}
			j++
		}
		if f.Len() != (i+1)-(j+1) {
			t.Fatal("Len of none empty fifo is invalid")
		}
	}
	for j < n {
		v, ok := f.Back().(int)
		if !ok {
			t.Fatal("Back retuned an invalid value, expected an int")
		} else if v != j {
			t.Fatalf("Back retuned %d instead of %d", v, j)
		}
		v, ok = f.PopBack().(int)
		if !ok {
			t.Fatal("PopBack retuned an invalid value, expected an int")
		} else if v != j {
			t.Fatalf("PopBack retuned %d instead of %d", v, j)
		}
		j++
	}
	testEmptyFifo(t, f)
}
