package piscine

func ListPushFront(l *List, data interface{}) {
	neNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = neNode
		l.Tail = neNode
		return
	}
	neNode.Next = l.Head
	l.Head = neNode
}
