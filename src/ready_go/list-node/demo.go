package main

import "fmt"

// Node 单链的数据结构
type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
}

// AddValue 添加成有序的链表
func (l *List) AddValue(value int) {
	if l.head == nil {
		node := Node{value: value}
		l.head = &node
		return
	}
	item := l.head
	for ; item.next != nil; item = item.next {
		if item.value == value {
			return
		}
		if item.value > value {
			tmpValue := item.value
			node := Node{value: tmpValue, next: item.next}
			item.value = value
			item.next = &node
			return
		}
	}
	// 处理最后的一个链表连接
	if item.value == value {
		return
	}
	node := Node{value: value}
	item.next = &node
}

// 删除链表节点的数据
func (l *List) deleteLink(value int) {
	if l.head == nil {
		return
	}
	item := l.head
	for ; item.next != nil; item = item.next {
		if item.value == value {
			item.value = item.next.value
			item.next = item.next.next
			break
		}
		if item.next.value == value && item.next.next == nil {
			item.next = nil
			break
		}
	}
}

// 翻转单链
func reserveLink(n *Node) *Node {
	if n == nil || n.next == nil {
		return n
	}
	// 这个是递归执行函数
	new := reserveLink(n.next)
	// 这里是从头节点开始下一个节点指向前一个节点
	n.next.next = n
	// 这里是把原来的节点指向置空，相当于实现了翻转
	n.next = nil
	return new
}

// 循环打印链表的每个值
func (l *List) printLink() {
	item := l.head
	if item != nil {
		for ; item.next != nil; item = item.next {
			fmt.Printf("next value %d\n", item.value)
		}
		fmt.Printf("end value %d\n", item.value)
	}
}

func NewOneLink() *List {
	return &List{head: nil}
}

func main() {
	nLink := NewOneLink()
	nLink.AddValue(1)
	nLink.AddValue(4)
	nLink.AddValue(5)
	nLink.AddValue(9)
	nLink.AddValue(9)
	nLink.AddValue(9)
	nLink.AddValue(2)

	nLink.printLink()
	fmt.Println("下面是删除节点之后----------------------")
	// 删除某个节点数据
	nLink.deleteLink(9)
	nLink.printLink()

	// 翻转单链
	nLink.head = reserveLink(nLink.head)
	fmt.Println("打印翻转之后----------------------")
	// 打印
	nLink.printLink()
}
