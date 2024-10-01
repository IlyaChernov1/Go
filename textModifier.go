package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Text struct {
	Content string
}

func (t *Text) textModifier() {
	// через цикл методом strings.Contains проверяем что в строке 2 пробела, с помощью метода Replace меняем 2 пробела на 1
	for strings.Contains(t.Content, "  ") {
		t.Content = strings.Replace(t.Content, "  ", " ", -1)
	}

	result := ""
	i := 0

	for i < len(t.Content) {

		if t.Content[i] == '-' { // Если текущий символ - минус
			if i > 0 && i < len(t.Content)-1 { // Проверяем есть ли символы слева и справа
				// Меняем местами символы
				result += string(t.Content[i+1]) + string(t.Content[i-1])
				i++ // Пропускаем следующий символ, так как он уже обработан
			}
		} else {
			result += string(t.Content[i]) // Добавляем текущий символ в результат
		}
		i++
	}

	t.Content = result

	fmt.Println(t.Content)
}

func main() {

	text := &Text{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите строку:")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
	}
}
