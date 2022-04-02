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

func TestDequeue(t *testing.T) {
	n2 := Node{20, nil}
	n1 := Node{10, &n2}
	queue := LinkedQueue{head: &n1, size: 2}
	got, _ := queue.Dequeue()
	want := 10

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}

}

func TestIsEmpty(t *testing.T) {
	queue := LinkedQueue{}
	got := queue.IsEmpty()
	if got != true {
		t.Errorf("got %t but want true", got)
	}
}

func TestLen(t *testing.T) {
	queue := LinkedQueue{}
	got := queue.Len()
	want := 0
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}

func TestFirst(t *testing.T) {
	node := Node{10, nil}
	queue := LinkedQueue{head: &node, size: 1}
	got, _ := queue.First()
	want := 10
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
