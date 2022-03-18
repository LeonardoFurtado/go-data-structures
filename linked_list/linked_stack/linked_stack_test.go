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
