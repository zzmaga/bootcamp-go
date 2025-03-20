package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListForEach applies a function f to each node of the list
func ListForEach(l *List, f func(*NodeL)) {
	if l == nil || l.Head == nil {
		return // Если список пустой, ничего не делаем
	}

	current := l.Head
	for current != nil {
		f(current)             // Применяем функцию к текущему узлу
		current = current.Next // Переходим к следующему узлу
	}
}

// Add2_node adds 2 to the data of the node if it's an int, or appends "2" if it's a string
func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

// Subtract3_node subtracts 3 from the data of the node if it's an int, or appends "-3" if it's a string
func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
