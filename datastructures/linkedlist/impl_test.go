package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AddNode_EmptyNode(t *testing.T) {
	ll := NewLinkedList()

	ll.AddNode(5)

	require.Equal(t, ll.head.value, 5)
	require.Equal(t, ll.tail.value, 5)
}

func Test_AddNode_TwoNodes(t *testing.T) {
	ll := NewLinkedList()

	ll.AddNode(1)
	ll.AddNode(2)

	require.Equal(t, ll.head.value, 1)
	require.Equal(t, ll.tail.value, 2)
}

func Test_AddNode_ThreeNodes(t *testing.T) {
	ll := NewLinkedList()

	ll.AddNode(1)
	ll.AddNode(2)
	ll.AddNode(3)

	require.Equal(t, ll.head.value, 1)
	require.Equal(t, ll.tail.value, 3)
	require.Equal(t, ll.head.next.value, 2)
	require.Equal(t, ll.tail.prev.value, 2)
}

func Test_GetNode_Empty(t *testing.T) {
	ll := NewLinkedList()

	n := ll.GetNode(1)

	require.Nil(t, n)
}

func Test_GetNode_SingleNode(t *testing.T) {
	ll := NewLinkedList()

	ll.AddNode(1)
	res := ll.GetNode(1)

	require.Equal(t, 1, res.value)
}

func Test_GetNode_DoubleNode(t *testing.T) {
	ll := NewLinkedList()

	ll.AddNode(1)
	ll.AddNode(2)
	res := ll.GetNode(2)

	require.Equal(t, 2, res.value)
}

func Test_GetNode_NoValue(t *testing.T) {
	ll := NewLinkedList()

	ll.AddNode(1)
	ll.AddNode(2)
	res := ll.GetNode(3)

	require.Nil(t, res)
}

func Test_RemoveNode_Empty(t *testing.T) {
	ll := NewLinkedList()

	n := ll.RemoveNode(1)

	require.Nil(t, n)
}

func Test_RemoveNode_SingleNode(t *testing.T) {
	ll := NewLinkedList()

	ll.AddNode(1)
	res := ll.RemoveNode(1)

	require.Equal(t, 1, res.value)
	require.Nil(t, ll.head)
	require.Nil(t, ll.tail)
}

func Test_RemoveNode_DoubleNode_LastNode(t *testing.T) {
	ll := NewLinkedList()

	ll.AddNode(1)
	ll.AddNode(2)
	res := ll.RemoveNode(2)

	require.Equal(t, 2, res.value)
	require.Equal(t, ll.head.value, 1)
	require.Equal(t, ll.tail.value, 1)
}

func Test_RemoveNode_DoubleNode_FirstNode(t *testing.T) {
	ll := NewLinkedList()

	ll.AddNode(1)
	ll.AddNode(2)
	res := ll.RemoveNode(1)

	require.Equal(t, 1, res.value)
	require.Equal(t, ll.head.value, 2)
	require.Equal(t, ll.tail.value, 2)
}

func Test_RemoveNode_TipleNode_MiddleNode(t *testing.T) {
	ll := NewLinkedList()

	ll.AddNode(1)
	ll.AddNode(2)
	ll.AddNode(3)
	res := ll.RemoveNode(2)

	require.Equal(t, 2, res.value)
	require.Equal(t, ll.head.value, 1)
	require.Equal(t, ll.tail.value, 3)
	require.Equal(t, ll.head.next.value, 3)
	require.Equal(t, ll.tail.prev.value, 1)
}

func Test_RemoveNode_TipleNode_NoneExistantValue(t *testing.T) {
	ll := NewLinkedList()

	ll.AddNode(1)
	ll.AddNode(2)
	ll.AddNode(3)
	res := ll.RemoveNode(4)

	require.Nil(t, res)
}
