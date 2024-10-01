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
	// замена любого кол-ва пробелов на 1
	for strings.Contains(t.Content, "  ") {
		t.Content = strings.Replace(t.Content, "  ", " ", -1)
	}

	//преобразовываем строку в срез рун чтобы верно выводилась латиница
	runes := []rune(t.Content)
	result := ""

	i := 0
	for i < len(runes) {

		if runes[i] == '-' { // Если текущий символ - минус
			if i > 0 && i < len(runes)-1 { // Проверяем есть ли символы слева и справа
				// Меняем местами символы
				result += string(runes[i+1]) + string(runes[i-1])
				i++ // Пропускаем следующий символ,т.к он уже обработан
			}
		} else {
			result += string(runes[i]) // Добавляем текущий символ в результат
		}
		i++
	}

	t.Content = result

	//меняем "+" на "!"
	t.Content = strings.ReplaceAll(t.Content, "+", "!")

	//удаление лишних пробелов
	t.Content = strings.TrimSpace(t.Content)

	//преобразуем строку в срез рун
	runes = []rune(t.Content)
	result = ""
	sum := 0

	for _, r := range runes {
		if r >= '0' && r <= '9' { //проверяем является ли символ цифрой
			sum += int(r - '0') //добавляем значение цифры к сумме
		} else {
			result += string(r) //добавляем цифру в результат
		}
	}
	result = strings.TrimSpace(result) //удаление пробелом в результате

	if sum > 0 {
		result += fmt.Sprintf(" %d", sum) //добавляем сумму цифр в конец строки
	}

	t.Content = result
	fmt.Println(t.Content)
}

func main() {
	text := &Text{}
	// Создаем новый сканер для чтения из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)

	// Просим пользователя ввести строку
	fmt.Println("Введите строку:")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
	}
}
