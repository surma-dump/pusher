package pusher

import (
	"testing"
)

type InterfaceImplementer1 struct {
	int
}

func (t *InterfaceImplementer1) Function() {
}

type InterfaceImplementer2 struct {
	float
}

func (t *InterfaceImplementer2) Function() {
}

type AnInterface interface {
	Function()
}

func TestPusher1(t *testing.T) {
	slice := make([]int, 0, 0)
	pusher := New(&slice)
	for i := 0; i < 100; i++ {
		pusher.Push(i)
	}
	for i := 0; i < 100; i++ {
		if slice[i] != i {
			t.Fatalf("Expected %d got %d\n", i, slice[i])
		}
	}
}

func TestPusher2(t *testing.T) {
	slice := make([]AnInterface, 0, 0)
	pusher := New(&slice)
	pusher.Push(InterfaceImplementer1{})
	pusher.Push(InterfaceImplementer2{})
}

