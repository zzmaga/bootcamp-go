package piscine

// ListReverse reverses the order of the elements in the linked list
func ListReverse(l *List) {
	var prev, current, next *NodeL
	current = l.Head

	// Iterate through the list and reverse the pointers
	for current != nil {
		next = current.Next // Save the next node
		current.Next = prev // Reverse the current node's Next pointer
		prev = current      // Move prev to the current node
		current = next      // Move to the next node
	}

	// After the loop, prev is the new head of the reversed list
	l.Tail = l.Head
	l.Head = prev
	// Tail should be the original head
}
