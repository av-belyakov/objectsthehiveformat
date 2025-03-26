package supportingfunctions

import (
	"strings"
	"unicode"
)

// GetWhitespace возвращает необходимое количество пробелов
func GetWhitespace(num int) string {
	var str string

	if num == 0 {
		return str
	}

	for i := 0; i < num; i++ {
		str += "  "
	}

	return str
}

// TrimIsNotLetter обрезает, с двух сторон, строку от символов не являющихся буквой или числом
func TrimIsNotLetter(str string) string {
	return strings.TrimFunc(str, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
