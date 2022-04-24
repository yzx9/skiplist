package skiplist

import (
	"fmt"
	"math/rand"
)

type SkipList[K Ordered, V any] struct {
	head  skipListNode[K, V]
	count int
	p     float32
}

type skipListNode[K Ordered, V any] struct {
	key   K
	val   *V
	level int
	prev  []*skipListNode[K, V]
	next  []*skipListNode[K, V]
}

var KeyNotExist = fmt.Errorf("key not exist")

func New[K Ordered, V any]() *SkipList[K, V] {
	return &SkipList[K, V]{
		head: skipListNode[K, V]{
			level: 1,
			next:  make([]*skipListNode[K, V], 1),
		},
		count: 0,
		p:     0.5,
	}
}

func (list *SkipList[K, V]) Insert(key K, val *V) {
	// search
	node := &list.head
	for i := node.level - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].key < key {
			node = node.next[i]
		}
	}

	// update value if exist
	if node.next[0] != nil && node.next[0].key == key {
		node.next[0].val = val
		return
	}

	// insert new node
	list.count++
	if list.head.level*list.head.level < list.count {
		list.head.level++
		list.head.next = append(list.head.next, nil)
	}

	level := list.randomLevel()
	newNode := &skipListNode[K, V]{
		key:   key,
		val:   val,
		level: level,
		prev:  make([]*skipListNode[K, V], level),
		next:  make([]*skipListNode[K, V], level),
	}
	for i := 0; i < level; i++ {
		for node.level-1 < i {
			node = node.prev[node.level-1]
		}

		next := node.next[i]
		newNode.prev[i] = node
		newNode.next[i] = next
		node.next[i] = newNode
		if next != nil {
			next.prev[i] = newNode
		}
	}
}

func (list *SkipList[K, V]) Delete(key K) error {
	// search
	node := &list.head
	for i := node.level - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].key < key {
			node = node.next[i]
		}
	}

	if node.next[0] == nil || node.next[0].key != key {
		return KeyNotExist
	}

	// delete
	list.count--
	node = node.next[0]
	for i := 0; i < node.level; i++ {
		node.prev[i].next[i] = node.next[i]
		if node.next[i] != nil {
			node.next[i].prev[i] = node.prev[i]
		}
	}
	return nil
}

func (list *SkipList[K, V]) Get(key K) (*V, error) {
	// search
	node := &list.head
	for i := node.level - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].key < key {
			node = node.next[i]
		}
	}

	if node.next[0] == nil || node.next[0].key != key {
		return nil, KeyNotExist
	}

	return node.next[0].val, nil
}

func (list SkipList[K, V]) Count() int { return list.count }

func (list SkipList[K, V]) randomLevel() int {
	newLevel := 1
	for newLevel < list.head.level && rand.Float32() < list.p {
		newLevel++
	}
	return newLevel
}
