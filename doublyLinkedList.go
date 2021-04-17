package ds

import (
	"fmt"
	"sync"
)

type DoublyLinkedList struct {
	Head  *Node
	size  int
	mutex sync.Mutex
}

func (d *DoublyLinkedList) incrementSize() {
	d.size += 1
}

func (d *DoublyLinkedList) decrementSize() {
	if d.size == 0 {
		return
	}
	d.size -= 1
}

func (d *DoublyLinkedList) Size() int {
	return d.size
}

func (d *DoublyLinkedList) Items() []Any {
	var items []Any

	for currentNode := d.Head; currentNode != nil; currentNode = currentNode.Next {
		items = append(items, currentNode.Value)
	}

	return items
}

func (d *DoublyLinkedList) Get(id int) (Any, error) {
	if d.size == 0 || id >= d.size {
		return nil, fmt.Errorf("index out of range")
	}

	currentNode := d.Head

	for i := 0; i < id; i++ {
		currentNode = currentNode.Next
	}

	return currentNode.Value, nil
}

func (d *DoublyLinkedList) Append(n *Node) {
	d.mutex.Lock()
	defer d.incrementSize()
	defer d.mutex.Unlock()

	if d.size == 0 {
		d.Head = n
		return
	}

	currentNode := d.Head

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	n.Previous = currentNode
	currentNode.Next = n
}

func (d *DoublyLinkedList) Prepend(n *Node) {
	d.mutex.Lock()
	defer d.incrementSize()
	defer d.mutex.Unlock()

	if d.size == 0 {
		d.Head = n
		return
	}

	n.Next = d.Head
	d.Head.Previous = n
	d.Head = n
}

func (d *DoublyLinkedList) Delete(id int) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if d.size == 0 || d.size <= id {
		return fmt.Errorf("index out of range")
	}

	defer d.decrementSize()
	if id == 0 {
		d.Head = d.Head.Next
		return nil
	}

	currentNode := d.Head
	var previousNode *Node

	for i := 0; i < id; i++ {
		previousNode = currentNode
		currentNode = currentNode.Next
	}

	// i == id
	previousNode.Next = currentNode.Next
	currentNode.Previous = previousNode

	return nil
}

func (d *DoublyLinkedList) Remove(value Any) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if d.size == 0 {
		return
	}

	for {
		if d.Head == nil {
			return
		}

		if d.Head.Value == value {
			d.Head = d.Head.Next
			d.decrementSize()
		} else {
			break
		}
	}

	currentNode := d.Head
	var previousNode *Node

	for currentNode != nil {
		if currentNode.Value == value {
			previousNode.Next = currentNode.Next
			d.decrementSize()
			if currentNode.Next != nil {
				currentNode.Next.Previous = previousNode
			}
		} else {
			previousNode = currentNode
		}
		currentNode = currentNode.Next
	}
}
