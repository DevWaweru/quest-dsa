package problemSolving

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkList struct {
	head *Node
}

func (ll *LinkList) insertBeginning(data interface{}) {
	node := &Node{data, ll.head}
	ll.head = node
}

func (ll *LinkList) insertEnd(data interface{}) {
	if ll.head == nil {
		ll.head = &Node{data, nil}
		return
	}
	obj := ll.head
	for obj.next != nil {
		obj = obj.next
	}
	obj.next = &Node{data, nil}
}

func (ll *LinkList) insertValues(values []interface{}) {
	for _, val := range values {
		fmt.Println(val)
		ll.insertEnd(val)
	}
}

func LinkedList() {
	linkedList := LinkList{}

	// insert at the beginning
	linkedList.insertBeginning("10")
	linkedList.insertEnd(20)
	linkedList.insertValues([]interface{}{1, 2, 3, 4})

	printLinkedList(linkedList.head)
}

func printLinkedList(node *Node) {
	for node != nil {
		fmt.Printf("%v, ", node.data)
		node = node.next
	}
	fmt.Println()
}
