package main

import (
	"fmt"
	"sync/atomic"
)

type DNode struct {
	data  interface{}
	index int
	next  *DNode
	prev  *DNode
}

func (n *DNode) GetData() interface{} {
	if n == nil {
		return nil
	}

	return n.data
}

func (n *DNode) GetNext() *DNode {
	return n.next
}

func (n *DNode) GetPrev() *DNode {
	return n.prev
}

func (n *DNode) String() string {
	if n == nil {
		return "nil"
	}

	return fmt.Sprintf("data[%d]: %v\nnext:\n%+v\n", n.index, n.data, n.next)
}

func NewDNode(data interface{}) *DNode {
	var n DNode
	n.data = data
	return &n
}

type DList struct {
	length int64
	head   *DNode
	tail   *DNode
}

func (l *DList) String() string {
	return fmt.Sprintf("length: %d,list:\n%+v", l.length, l.head)
}

func NewDList() *DList {
	var l DList
	return &l
}

func (l *DList) Len() int {
	return int(l.length)
}

func (l *DList) Append(data interface{}) *DList {
	defer atomic.AddInt64(&l.length, 1)

	if l.head == nil {
		var node = NewDNode(data)
		node.index = int(l.length)
		l.head = node
		l.tail = node
		return l
	}

	var tail = NewDNode(data)
	tail.index = int(l.length)
	l.tail.next = tail
	tail.prev = l.tail
	l.tail = tail
	return l
}

func (l *Dlist) Index(index int) *DNode {
	if index > l.length-1 {
		return nil
	}

	var node = l.head
	for node.index != index {
		node = node.next
	}

	return node
}

func (l *DList) Restore() {
	if l.length == 0 {
		return
	}

	var node = l.head
	var index int

	for node.next != nil {
		node.index = index
		node = node.next
		index++
	}

	l.tail = node
	return
}

func (l *DList) CutAfter(index int) (*DList, *DList) {
	if index == l.length-1 {
		return l, nil
	}

	var tail = l.Index(index)
	if tail == nil {
		return l, nil
	}

	var newHead = tail.next
	tail.next = nil
	l.length = index + 1

	newHead.prev = nil
	var newList = NewDList()
	newList.head = newHead
	newList.length = l.length - index - 1
	newList.Resotre()

	return l, newList
}

func (l *DList) CutBefore(index int) (*DList, *DList) {
	if index == 0 {
		return nil, l
	}

}

func (l *DList) Extend(list *DList) *DList {
	l.tail.next = list.head
	list.head.prev = l.tail
	l.Restore()
	return l
}

func (l *DList) InsertBefore(data interface{}, index int) {
	if index == l.length-1 {
		return l.Append(data)
	}

	var headList, tailList = l.CutAt(index)
}
