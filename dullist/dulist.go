package dullist

import (
	"fmt"

	"github.com/yzx9/skiplist"
)

type DulList[K skiplist.Ordered, V any] struct {
	head *dulListNode[K, V]
}

type dulListNode[K skiplist.Ordered, V any] struct {
	key  K
	val  *V
	prev *dulListNode[K, V]
	next *dulListNode[K, V]
}

var KeyNotExist = fmt.Errorf("key not exist")

func New[K skiplist.Ordered, V any]() *DulList[K, V] {
	return &DulList[K, V]{head: nil}
}

func (l *DulList[K, V]) Insert(key K, val *V) {
	if l.head == nil {
		l.head = &dulListNode[K, V]{
			key:  key,
			val:  val,
			prev: nil,
			next: nil,
		}
		return
	}

	if l.head.key > key {
		l.head = &dulListNode[K, V]{
			key:  key,
			val:  val,
			prev: nil,
			next: l.head,
		}
		l.head.next.prev = l.head
		return
	}

	node := l.head
	for node.next != nil && node.next.key < key {
		node = node.next
	}

	node.next = &dulListNode[K, V]{
		key:  key,
		val:  val,
		prev: node,
		next: node.next,
	}
	if node.next.next != nil {
		node.next.next.prev = node.next
	}
	return
}

func (l *DulList[K, V]) Delete(key K) error {
	if l.head == nil || l.head.key > key {
		return KeyNotExist
	}

	if l.head.key == key {
		l.head = l.head.next
		return nil
	}

	node := l.head
	for node.next != nil && node.next.key < key {
		node = node.next
	}

	if node.next == nil || node.next.key != key {
		return KeyNotExist
	}

	node.next = node.next.next
	return nil
}

func (l *DulList[K, V]) Get(key K) (*V, error) {
	if l.head == nil || l.head.key > key {
		return nil, KeyNotExist
	}

	if l.head.key == key {
		return l.head.val, nil
	}

	node := l.head
	for node.next != nil && node.next.key < key {
		node = node.next
	}

	if node.next == nil || node.next.key != key {
		return nil, KeyNotExist
	}

	return node.next.val, nil
}
