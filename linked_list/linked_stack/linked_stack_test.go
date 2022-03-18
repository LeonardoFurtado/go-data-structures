package linked_stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	n3 := Node{10, nil}
	n2 := Node{20, &n3}
	stack := LinkedStack{&n2, 2}
	stack.Push(30)
	if stack.head.element != 30 {
		t.Errorf("Value was not inserted on top.")
	}
}

func TestPop(t *testing.T) {
	n3 := Node{10, nil}
	n2 := Node{20, &n3}
	n1 := Node{30, &n2}
	stack := LinkedStack{&n1, 3}
	got := stack.Pop()
	want := 30
	if want != got {
		t.Errorf("Want %d but got %d", want, got)
	}

	if stack.size != 2 {
		t.Errorf("The stack size was not decreased")
	}
}

func TestTop(t *testing.T) {
	n3 := Node{10, nil}
	n2 := Node{20, &n3}
	n1 := Node{30, &n2}
	stack := LinkedStack{&n1, 3}
	got := stack.Top()
	want := 30

	if got != want {
		t.Errorf("want %d but got %d", want, got)
	}
}
