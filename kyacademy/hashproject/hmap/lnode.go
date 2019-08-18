package hmap

import "fmt"

//
type KV struct {
	Key   string
	Value string
}

//
type Node struct {
	Data KV
	Next *Node
}

// 创建头结点
func CreateHead(data KV) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

// 添加节点
func Add(data KV, node *Node) *Node {
	newNode := &Node{
		Data: data,
		Next: nil,
	}
	node.Next = newNode

	return newNode
}

// 遍历
func ShowNodes(head *Node) {
	node := head
	for node.Next != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
	fmt.Println(node.Data)
}

//
func Tail(head *Node) *Node {
	node := head
	for node.Next != nil {
		node = node.Next
	}

	return node
}

//
func FindValueByKey(key string, head *Node) string {
	node := head
	for node != nil {
		if node.Data.Key == key {
			return node.Data.Value
		}
		node = node.Next
	}

	return ""
}
