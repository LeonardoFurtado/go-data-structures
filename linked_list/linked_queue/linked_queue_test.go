package linked_queue

import "testing"

func TestEnqueue(t *testing.T) {
	n3 := Node{10, nil}
	n2 := Node{10, &n3}
	n1 := Node{10, &n2}
	queue := LinkedQueue{&n1, &n3, 3}
	queue.Enqueue(40)
	if queue.tail.element != 40 {
		t.Errorf("Value was not enqueued")
	}
}

func TestIsEmpty(t *testing.T) {
	queue := LinkedQueue{}
	got := queue.IsEmpty()
	if got != true {
		t.Errorf("got false but want true")
	}
}
