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
	prevNodes := make([]*skipListNode[K, V], list.head.level)
	for i := list.head.level - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].key < key {
			node = node.next[i]
		}
		prevNodes[i] = node
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
		next:  make([]*skipListNode[K, V], level),
	}
	for i := 0; i < level && i < len(prevNodes); i++ {
		newNode.next[i] = prevNodes[i].next[i]
		prevNodes[i].next[i] = newNode
	}
}

func (list *SkipList[K, V]) Delete(key K) error {
	// search
	node := &list.head
	prevNodes := make([]*skipListNode[K, V], list.head.level)
	for i := list.head.level - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].key < key {
			node = node.next[i]
		}
		prevNodes[i] = node
	}

	if node.next[0] == nil || node.next[0].key != key {
		return KeyNotExist
	}

	// delete
	list.count--
	node = node.next[0]
	for i := 0; i < node.level; i++ {
		prevNodes[i].next[i] = node.next[i]
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
