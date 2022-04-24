package dullist

import (
	"fmt"

	"github.com/yzx9/skiplist"
)

type DulList[K skiplist.Ordered, V any] struct {
	head  dulListNode[K, V]
	count int
}

type dulListNode[K skiplist.Ordered, V any] struct {
	key  K
	val  *V
	prev *dulListNode[K, V]
	next *dulListNode[K, V]
}

var KeyNotExist = fmt.Errorf("key not exist")

func New[K skiplist.Ordered, V any]() *DulList[K, V] {
	return &DulList[K, V]{
		head: dulListNode[K, V]{},
	}
}

func (l *DulList[K, V]) Insert(key K, val *V) {
	// search
	node := &l.head
	for node.next != nil && node.next.key < key {
		node = node.next
	}

	// update value if exist
	if node.next != nil && node.next.key == key {
		node.next.val = val
		return
	}

	// insert new node
	l.count++
	next := node.next
	node.next = &dulListNode[K, V]{
		key:  key,
		val:  val,
		prev: node,
		next: next,
	}
	if next != nil {
		next.prev = node.next
	}
}

func (l *DulList[K, V]) Delete(key K) error {
	// search
	node := &l.head
	for node.next != nil && node.next.key < key {
		node = node.next
	}

	if node.next == nil || node.next.key != key {
		return KeyNotExist
	}

	// delete
	l.count--
	node.next = node.next.next
	if node.next != nil {
		node.next.prev = node
	}
	return nil
}

func (l *DulList[K, V]) Get(key K) (*V, error) {
	//search
	node := &l.head
	for node.next != nil && node.next.key < key {
		node = node.next
	}

	if node.next == nil || node.next.key != key {
		return nil, KeyNotExist
	}

	return node.next.val, nil
}

func (l DulList[K, V]) Count() int { return l.count }
