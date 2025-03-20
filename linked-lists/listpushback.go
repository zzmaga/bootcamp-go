package piscine

func ListPushBack(l *List, data interface{}) {
	// Создаём новый узел
	newNode := &NodeL{Data: data, Next: nil}

	// Если список пустой
	if l.Head == nil && l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	// Если список не пустой
	l.Tail.Next = newNode
	l.Tail = newNode
}
