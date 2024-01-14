package problemSolving

// Problem:
// Given head, the head of a linked list, determine if the linked list has a cycle in it.
// There is a cycle in a linked list if there is some node in the list that can be reached
// again by continuously following the next pointer. Internally, pos is used to denote the
// index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
// Return true if there is a cycle in the linked list. Otherwise, return false.
// head = [3,2,0,-4], pos = 1 ==> true
func linkedListCycle(head *LinkList) bool {
	slowPointer := head.head
	fastPointer := head.head
	for fastPointer != nil && fastPointer.next != nil {
		slowPointer = slowPointer.next
		fastPointer = fastPointer.next.next
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}

func happyNumber() bool {

	return true
}

func FastSlowPointers() {
	linked := LinkList{}
	linked.insertValues([]interface{}{3, 2, 0, -4})
	linkedListCycle(&linked)
	happyNumber()
}
