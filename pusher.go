// A simple replacement for container/vector
// A Pusher works on an already existing, arbitrarily typed
// slice and provied a function to push new values. The slice
// will be resized if it runs out of storage.
package pusher

import (
	. "reflect"
)

// The Pusher struct provied the Push() function
// to push new values into the struct.
type Pusher struct {
	data      *SliceValue
	Push      PushFunc
}

// This is the signature of a Push() function.
// It appends the passed value to the end of the array
// if the type assertions hold.
type PushFunc func(obj interface{})

// Create a new pusher which works on an existing slice.
// v has to be a pointer to slice
func New(v interface{}) (p *Pusher) {
	p = new(Pusher)
	p.data = getSliceValue(v)
	p.Push = func(obj interface{}) {
		pusherFunc(p, obj)
	}
	return p
}

func getSliceValue(v interface{}) *SliceValue {
	return NewValue(v).(*PtrValue).Elem().(*SliceValue)
}

func isEqualType(p *Pusher, obj interface{}) bool {
	return Typeof(obj) == p.getSliceType().Elem()
}

func (p *Pusher) getSliceType() *SliceType {
	return p.data.Type().(*SliceType)
}

func pusherFunc(p *Pusher, obj interface{}) {
	if !isEqualType(p, obj) {
		panic("Incompatible types")
	}
	len, cap := p.data.Len(), p.data.Cap()
	if len == cap {
		newData := MakeSlice(p.getSliceType(), len, (len+10)*2)
		ArrayCopy(newData, p.data)
		p.data.SetValue(newData)
	}
	p.data.SetLen(len+1)
	p.data.Elem(len).SetValue(NewValue(obj))
	return
}