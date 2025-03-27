package objectsthehiveformat

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/casedetails"
	"github.com/av-belyakov/objectsthehiveformat/caseobject"
	"github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewTypeEventForCase создаёт новый объект типа TypeEventForCase
func NewTypeEventForCase() *TypeEventForCase {
	return &TypeEventForCase{
		CommonEventType: *common.NewCommonEventType(),
		Details:         *casedetails.NewEventCaseDetails(),
		Object:          *caseobject.NewEventCaseObject(),
	}
}

func (e *TypeEventForCase) Get() *TypeEventForCase {
	return e
}

func (e *TypeEventForCase) GetDetails() casedetails.EventCaseDetails {
	return e.Details
}

// SetValueDetails устанавливает значение типа EventDetails для поля Details
func (e *TypeEventForCase) SetValueDetails(v casedetails.EventCaseDetails) {
	e.Details = v
}

func (e *TypeEventForCase) GetObject() caseobject.EventCaseObject {
	return e.Object
}

// SetValueObject устанавливает значение типа EventForEsCaseObject для поля Object
func (e *TypeEventForCase) SetValueObject(v caseobject.EventCaseObject) {
	e.Object = v
}

// ToStringBeautiful форматированный вывод
func (em TypeEventForCase) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(em.CommonEventType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'details':\n", ws))
	str.WriteString(em.Details.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%s'object':\n", ws))
	str.WriteString(em.Object.ToStringBeautiful(num + 1))

	return str.String()
}
