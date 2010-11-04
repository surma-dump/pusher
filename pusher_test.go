package pusher

import (
	"testing"
)

type InterfaceImplementer1 struct {
	a int
}

func (t *InterfaceImplementer1) aFunction() {
}

type InterfaceImplementer2 struct {
	a float
}

func (t *InterfaceImplementer2) aFunction() {
}

type AnInterface interface {
	aFunction()
}

func error(t *testing.T) {
	if x := recover(); x != nil {
		t.Fatalf("Panic: %s\n", x)
	}
}

func TestPusher1(t *testing.T) {
	defer error(t)
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

func createElement(i int) AnInterface {
	if i % 2 == 0 {
		return &InterfaceImplementer1{}
	}
	return &InterfaceImplementer2{}
}

func TestPusher2(t *testing.T) {
	defer error(t)
	slice := make([]AnInterface, 0, 0)
	pusher := New(&slice)
	ok := false
	for i := 0; i < 99; i ++ {
		pusher.Push(createElement(i))
	}
	for i := 0; i < 99; i ++ {
		if i % 2 == 0 {
			_, ok = slice[i].(*InterfaceImplementer1)
		} else {
			_, ok = slice[i].(*InterfaceImplementer2)
		}
		if !ok {
			t.Fatalf("Could not get Interface-Types from slice")
		}

	}

}
