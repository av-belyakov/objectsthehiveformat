package objectsthehiveformat

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/alertdetails"
	"github.com/av-belyakov/objectsthehiveformat/alertobject"
	"github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewTypeEventForAlert создаёт новый объект типа TypeEventForAlert
func NewTypeEventForAlert() *TypeEventForAlert {
	return &TypeEventForAlert{
		CommonEventType: *common.NewCommonEventType(),
		Details:         *alertdetails.NewEventAlertDetails(),
		Object:          *alertobject.NewEventAlertObject(),
	}
}

// Get возвращает объект типа EventMessageForEsAlert
func (e *TypeEventForAlert) Get() *TypeEventForAlert {
	return e
}

func (e *TypeEventForAlert) GetDetails() alertdetails.EventAlertDetails {
	return e.Details
}

// SetValueDetails значение типа alertdetails.EventAlertDetails для поля Details
func (e *TypeEventForAlert) SetValueDetails(v alertdetails.EventAlertDetails) {
	e.Details = v
}

func (e *TypeEventForAlert) GetObject() alertobject.EventAlertObject {
	return e.Object
}

// SetValueObject значение типа alertobject.EventAlertObject для поля Object
func (e *TypeEventForAlert) SetValueObject(v alertobject.EventAlertObject) {
	e.Object = v
}

// ToStringBeautiful форматированный вывод
func (e *TypeEventForAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(e.CommonEventType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'details':\n", ws))
	str.WriteString(e.Details.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%s'object':\n", ws))
	str.WriteString(e.Object.ToStringBeautiful(num + 1))

	return str.String()
}
