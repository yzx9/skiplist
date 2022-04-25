package linklist

import (
	"fmt"

	"github.com/yzx9/skiplist"
)

type LinkList[K skiplist.Ordered, V any] struct {
	head  linkListNode[K, V]
	count int
}

type linkListNode[K skiplist.Ordered, V any] struct {
	key  K
	val  *V
	next *linkListNode[K, V]
}

var KeyNotExist = fmt.Errorf("key not exist")

func New[K skiplist.Ordered, V any]() *LinkList[K, V] {
	return &LinkList[K, V]{
		head: linkListNode[K, V]{},
	}
}

func (l *LinkList[K, V]) Insert(key K, val *V) {
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
	node.next = &linkListNode[K, V]{
		key:  key,
		val:  val,
		next: next,
	}
}

func (l *LinkList[K, V]) Delete(key K) error {
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
	return nil
}

func (l *LinkList[K, V]) Get(key K) (*V, error) {
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

func (l LinkList[K, V]) Count() int { return l.count }
