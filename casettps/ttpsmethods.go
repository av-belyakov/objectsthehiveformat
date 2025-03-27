package casettps

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewTtps создаёт объект типа Ttps
func NewTtps() *Ttps {
	return &Ttps{
		Ttp: []Ttp(nil),
	}
}

func (ttps *Ttps) SetTtps(list []Ttp) {
	ttps.Ttp = list
}

func (ttps *Ttps) GetTtps() []Ttp {
	return ttps.Ttp
}

func (ttps *Ttps) Set(v Ttp) {
	ttps.Ttp = append(ttps.Ttp, v)
}

// ToStringBeautiful форматированный вывод
func (tm Ttps) ToStringBeautiful(num int) string {
	return fmt.Sprintf("%s'ttp': \n%s", supportingfunctions.GetWhitespace(num), func(l []Ttp) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+1), k+1))
			str.WriteString(v.ToStringBeautiful(num + 2))
		}

		return str.String()
	}(tm.Ttp))
}
