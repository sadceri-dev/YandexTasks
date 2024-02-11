package main

import (
	"fmt"
	"regexp"
	"strings"
)

func normalizePhoneNumber(number string) string {
	// Удаляем все символы, кроме цифр
	re := regexp.MustCompile(`\D`)
	number = re.ReplaceAllString(number, "")
	// Добавляем префикс "8495", если номер состоит из 7 цифр (автоматически добавляется код 495)
	if len(number) == 7 {
		number = "8495" + number
	} else if len(number) == 10 {
		// Для номеров, начинающихся на 7 или 8 и имеющих 10 цифр, предполагаем, что это внутренний формат и добавляем "8"
		number = "8" + number
	} else if strings.HasPrefix(number, "7") && len(number) == 11 {
		// Если номер начинается с 7 и состоит из 11 цифр, заменяем 7 на 8
		number = "8" + number[1:]
	}
	return number
}

func main() {
	var newNumber string
	fmt.Scanln(&newNumber)

	phoneBook := make([]string, 3)
	for i := range phoneBook {
		fmt.Scanln(&phoneBook[i])
	}

	normalizedNewNumber := normalizePhoneNumber(newNumber)
	for _, number := range phoneBook {
		normalizedNumber := normalizePhoneNumber(number)
		if normalizedNumber == normalizedNewNumber {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
