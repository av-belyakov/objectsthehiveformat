package supportingfunctions

import (
	"fmt"
	"strings"
)

func ToStringBeautifulSlice[T any](num int, l []T) string {
	str := strings.Builder{}
	ws := GetWhitespace(num + 1)

	for k, v := range l {
		str.WriteString(fmt.Sprintf("%s%d. '%v'\n", ws, k+1, v))
	}

	return str.String()
}

func ToStringBeautifulMapSlice(num int, m map[string][]string) string {
	str := strings.Builder{}

	for k, v := range m {
		str.WriteString(fmt.Sprintf("%s'%s'\n", GetWhitespace(num+1), k))
		for key, value := range v {
			str.WriteString(fmt.Sprintf("%s%d. %s\n", GetWhitespace(num+2), key+1, value))
		}
	}

	return str.String()
}
