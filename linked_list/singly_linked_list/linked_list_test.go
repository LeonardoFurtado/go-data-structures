package singly_linked_list

import (
	"testing"
)

// AddLast(t) adds an Item t to the end of the linked list     -
// Insert(i, t) adds an Item t at position i                  -
// RemoveAt(i) removes a node at position i				      -
// IndexOf() returns the position of the Item t               -
// IsEmpty() returns true if the list is empty                -
// Size() returns the linked list size                        -
// String() returns a string representation of the list
// Head() returns the first node, so we can iterate on it

//TODO: Refactor

func TestAddFirst(t *testing.T) {
	n3 := Node{30, nil}
	n2 := Node{20, &n3}
	n1 := Node{10, &n2}
	list := LinkedList{&n1, 3}
	list.AddFirst(5)

	if list.head.value != 5 {
		t.Errorf("Value not inserted at head.")
	}
}

func TestAddLast(t *testing.T) {
	list := LinkedList{}
	list.AddLast(5)
	list.AddLast(10)
	list.AddLast(15)

	if list.head.value != 5 {
		t.Errorf("Value not inserted!")
	}
	if list.head.next.value != 10 {
		t.Errorf("Value not inserted!")
	}
	if list.head.next.next.value != 15 {
		t.Errorf("Value not inserted!")
	}
}

func TestInsert(t *testing.T) {
	n2 := Node{20, nil}
	n1 := Node{10, &n2}
	list := LinkedList{&n1, 2}
	list.Insert(2, 15)

	if list.head.value != 10 || list.head.next.value != 15 || list.head.next.next.value != 20 {
		t.Errorf("Value not inserted correctly")
	}
}

func TestRemoveFirst(t *testing.T) {
	n3 := Node{30, nil}
	n2 := Node{20, &n3}
	n1 := Node{10, &n2}
	list := LinkedList{&n1, 3}
	list.RemoveFirst()

	if list.head.value != 20 {
		t.Errorf("Value was not removed from head")
	}
}

func TestRemoveAt(t *testing.T) {
	t.Run("Testing remove first element", func(t *testing.T) {
		n3 := Node{30, nil}
		n2 := Node{20, &n3}
		n1 := Node{10, &n2}
		list := LinkedList{&n1, 3}
		list.RemoveAt(1)

		if list.head != &n2 {
			t.Errorf("Value not removed from list! list head should be %d, but is %v", list.head.value, n2.value)
		}
	})

	t.Run("Testing remove midle element", func(t *testing.T) {
		n3 := Node{30, nil}
		n2 := Node{20, &n3}
		n1 := Node{10, &n2}
		list := LinkedList{&n1, 3}
		list.RemoveAt(2)

		if list.head.next.value != n3.value {
			t.Errorf("Value not removed from list! list head next should be %d, but is %d", n3.value, list.head.next.value)
		}
	})

	t.Run("Testing remove final element", func(t *testing.T) {
		n3 := Node{30, nil}
		n2 := Node{20, &n3}
		n1 := Node{10, &n2}
		list := LinkedList{&n1, 3}
		list.RemoveAt(3)

		if list.head.next.next != nil {
			t.Errorf("Value not removed from list!")
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Testing empty LinkedList", func(t *testing.T) {
		list := LinkedList{}

		if list.IsEmpty() != true {
			t.Errorf("List is not empty")
		}
	})
}

func TestIndexOf(t *testing.T) {
	n3 := Node{30, nil}
	n2 := Node{20, &n3}
	n1 := Node{10, &n2}
	list := LinkedList{&n1, 3}
	got := list.IndexOf(30)
	expected := 3

	if got != expected {
		t.Errorf("got %d but expected %d", got, expected)
	}
}

func TestSize(t *testing.T) {
	n3 := Node{30, nil}
	n2 := Node{20, &n3}
	n1 := Node{10, &n2}
	list := LinkedList{&n1, 3}
	got := list.Size()
	expected := 3

	if got != expected {
		t.Errorf("got %d but expected %d", got, expected)
	}

}
