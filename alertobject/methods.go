package alertobject

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewEventAlertObject создаёт новый объект типа EventAlertObject
func NewEventAlertObject() *EventAlertObject {
	return &EventAlertObject{
		CommonEventAlertObject: *common.NewCommonEventAlertObject(),
		Tags:                   make(map[string][]string),
		TagsAll:                []string(nil),
		CustomFields:           common.CustomFields{},
	}
}

func (e *EventAlertObject) Get() *EventAlertObject {
	return e
}

func (e *EventAlertObject) GetTags() map[string][]string {
	return e.Tags
}

// SetValueTags значение в Tags по ключу
func (e *EventAlertObject) SetValueTags(k, v string) bool {
	if _, ok := e.Tags[k]; !ok {
		e.Tags[k] = []string(nil)
	}

	v = supportingfunctions.TrimIsNotLetter(v)
	for _, value := range e.Tags[k] {
		if v == value {
			return false
		}
	}

	e.Tags[k] = append(e.Tags[k], v)

	return true
}

// SetAnyTags значение для поля Tags
func (e *EventAlertObject) SetAnyTags(k string, i interface{}) bool {
	return e.SetValueTags(k, fmt.Sprint(i))
}

func (a *EventAlertObject) GetTagsAll() []string {
	return a.TagsAll
}

// SetValueTagsAll значение в список поля TagsAll
func (a *EventAlertObject) SetValueTagsAll(v string) {
	a.TagsAll = append(a.TagsAll, v)
}

// SetAnyTagsAll значение в список поля TagsAll
func (a *EventAlertObject) SetAnyTagsAll(i interface{}) {
	a.TagsAll = append(a.TagsAll, fmt.Sprint(i))
}

func (e *EventAlertObject) GetCustomFields() common.CustomFields {
	return e.CustomFields
}

// SetValueCustomFields значение для поля CustomFields
func (e *EventAlertObject) SetValueCustomFields(v common.CustomFields) {
	e.CustomFields = v
}

// ToStringBeautiful форматированный вывод
func (e *EventAlertObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(e.CommonEventAlertObject.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, supportingfunctions.ToStringBeautifulMapSlice(num, e.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, e.TagsAll)))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, common.CustomFieldsToStringBeautiful(e.CustomFields, num)))

	return str.String()
}
