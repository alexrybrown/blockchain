package linkedlist

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) AddNode(v int) {
	n := NewNode(v)

	if ll.head == nil {
		ll.head = n
		ll.tail = n
		return
	}

	cur := ll.head
	nextNode := cur.next
	for nextNode != nil {
		cur = nextNode
		nextNode = cur.next
	}

	cur.next = n
	n.prev = cur
	ll.tail = n
}

func (ll *LinkedList) RemoveNode(v int) *Node {
	if ll.head == nil {
		return nil
	}

	cur := ll.head
	nextNode := cur.next
	for nextNode != nil && cur.value != v {
		cur = nextNode
		nextNode = cur.next
	}

	if cur.value != v {
		return nil
	}

	if cur == ll.head {
		if ll.head.next != nil {
			ll.head = ll.head.next
			ll.head.prev = nil
			return cur
		}

		ll.head = nil
		ll.tail = nil
		return cur
	}

	if cur.next == nil {
		ll.tail = cur.prev
		ll.tail.next = nil
		return cur
	}

	cur.prev.next = cur.next
	cur.next.prev = cur.prev
	return cur
}

func (ll *LinkedList) GetNode(v int) *Node {
	if ll.head == nil {
		return nil
	}

	cur := ll.head
	nextNode := cur.next
	for nextNode != nil && cur.value != v {
		cur = nextNode
		nextNode = cur.next
	}

	if cur.value != v {
		return nil
	}

	return cur
}
