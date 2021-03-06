package ds

import (
	// "sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSLLAppend(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}

	elements := []Any{1, 9, 9, 8}
	seedList(&list, elements)

	assert.Equal(len(elements), list.Size())

	items := list.Items()

	for i := 0; i < len(items); i++ {
		assert.Equal(elements[i], items[i])
	}
}

func TestSLLDelete(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}
	elements := []Any{1, 2, 3, 4}

	seedList(&list, elements)
	assert.NotNil(list.Delete(5), "want out of range error; got nil")

	testFn := func(list *SinglyLinkedList, elements []Any) {
		assert.Equal(len(elements), list.Size)

		items := list.Items()
		for i := 0; i < list.Size(); i++ {
			assert.Equal(elements[i], items[i])
		}
	}

	list.Delete(0)
	elements = elements[1:]
	testFn(&list, elements)

	list.Delete(1)
	elements = []Any{2, 4}
	testFn(&list, elements)
}

func TestSLLRemove(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}
	elements := []Any{1, 3, 3, 4}
	seedList(&list, elements)

	list.Remove(3)
	elements = []Any{1, 4}
	assert.Equal(len(elements), list.Size())

	items := list.Items()
	for i := 0; i < list.Size(); i++ {
		assert.Equal(elements[i], items[i])
	}

	list = SinglyLinkedList{}
	elements = []Any{3, 3, 3, 4}
	for _, elem := range elements {
		list.Append(&Node{Value: elem})
	}

	list.Remove(3)
	assert.Equal(1, list.Size())
	assert.Equal(4, list.Head.Value)
}

func TestSSLPrepend(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}

	testFn := func(list *SinglyLinkedList, size, value int) {
		assert.Equal(size, list.Size())

		assert.Equal(value, list.Head.Value)
	}

	list.Prepend(&Node{Value: 1})
	testFn(&list, 1, 1)

	list.Prepend(&Node{Value: 2})
	testFn(&list, 2, 2)
}

func TestSSLReverse(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}
	elements := []Any{1, 2, 3, 4}
	seedList(&list, elements)

	list.Reverse()
	assert.Equal(4, list.Size())

	items := list.Items()
	elements = []Any{4, 3, 2, 1}
	for i := 0; i < list.Size(); i++ {
		assert.Equal(elements[i], items[i])
	}
}

func TestSSLClear(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}
	seedList(&list, []Any{1, 2, 3, 4})
	list.Clear()

	assert.Equal(0, list.Size())
	assert.Nil(list.Head, "list.Head still points to a node")
}

func TestSLLUnique(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}
	elements := []Any{3, 3, 1, 8, 0, 3, 8}
	seedList(&list, elements)
	list.Unique()

	elements = []Any{3, 1, 8, 0}
	assert.Equal(len(elements), list.Size())

	items := list.Items()
	for i := 0; i < list.Size(); i++ {
		assert.Equal(elements[i], items[i])
	}
}

func TestSLLSwap(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}
	elements := []Any{1, 2, 3, 4}
	seedList(&list, elements)

	// test valid swap
	list.Swap(1, 3)
	assert.Equal(len(elements), list.Size())

	items := list.Items()
	elements = []Any{1, 4, 3, 2}

	for i := 0; i < list.Size(); i++ {
		assert.Equal(elements[i], items[i])
	}

	// test swap head with tail
	list.Swap(0, 3)
	elements = []Any{2, 4, 3, 1}
	assert.Equal(len(elements), list.Size())

	items = list.Items()
	for i := 0; i < list.Size(); i++ {
		assert.Equal(elements[i], items[i])
	}

	// test out of range indices
	assert.Error(list.Swap(10, 20))

	// test swap where 1st index > 2nd index
	list.Swap(3, 1)
	assert.Equal(len(elements), list.Size())

	items = list.Items()
	elements = []Any{2, 1, 3, 4}

	for i := 0; i < list.Size(); i++ {
		assert.Equal(elements[i], items[i])
	}
}

func TestSLLGet(t *testing.T) {
	assert := assert.New(t)
	var list SinglyLinkedList

	item, err := list.Get(0)
	assert.Nil(item)
	assert.Error(err)

	elements := []Any{1, 9, 9, 8}
	seedList(&list, elements)

	for i := 0; i < len(elements); i++ {
		item, err = list.Get(i)
		assert.Nil(err)
		assert.Equal(elements[i], item)
	}

	item, err = list.Get(len(elements))
	assert.Nil(item)
	assert.Error(err)
}
