package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	// Проверка на невалидные входные данные
	if l == nil || pos < 0 {
		return nil
	}

	current := l // текущий узел
	index := 0   // текущая позиция

	// Перебор списка до достижения нужной позиции
	for current != nil {
		if index == pos {
			return current // Найден нужный узел
		}
		current = current.Next // Переход к следующему узлу
		index++                // Увеличение индекса
	}

	// Если позиция не найдена
	return nil
}
