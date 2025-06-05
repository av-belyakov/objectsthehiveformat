package caseobject

import (
	"fmt"
	"strings"

	"slices"

	"github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

func NewEventCaseObject() *EventCaseObject {
	return &EventCaseObject{
		CommonEventCaseObject: *common.NewCommonEventCaseObject(),
		Tags:                  make(map[string][]string),
		TagsAll:               []string(nil),
		CustomFields:          common.CustomFields{},
	}
}

func (e *EventCaseObject) Get() *EventCaseObject {
	return e
}

func (e *EventCaseObject) GetTags() map[string][]string {
	return e.Tags
}

// SetValueTags добаляет значение в Tags по ключу
func (e *EventCaseObject) SetValueTags(k, v string) bool {
	if e.Tags == nil {
		e.Tags = make(map[string][]string)
	}

	if _, ok := e.Tags[k]; !ok {
		e.Tags[k] = []string(nil)
	}

	v = supportingfunctions.TrimIsNotLetter(v)
	if slices.Contains(e.Tags[k], v) {
		return false
	}

	e.Tags[k] = append(e.Tags[k], v)

	return true
}

// SetAnyTags устанавливает ЛЮБОЕ значение для поля Tags
func (e *EventCaseObject) SetAnyTags(k string, i any) bool {
	return e.SetValueTags(k, fmt.Sprint(i))
}

func (a *EventCaseObject) GetTagsAll() []string {
	return a.TagsAll
}

// SetValueTagsAll добавляет значение STRING в список поля TagsAll
func (a *EventCaseObject) SetValueTagsAll(v string) {
	if a.TagsAll == nil {
		a.TagsAll = []string(nil)
	}

	a.TagsAll = append(a.TagsAll, v)
}

// SetAnyTagsAll добавляет ЛЮБОЕ значение в список поля TagsAll
func (a *EventCaseObject) SetAnyTagsAll(i any) {
	a.SetValueTagsAll(fmt.Sprint(i))
}

func (e *EventCaseObject) GetCustomFields() common.CustomFields {
	return e.CustomFields
}

// SetValueCustomFields устанавливает STRING значение для поля CustomFields
func (e *EventCaseObject) SetValueCustomFields(v common.CustomFields) {
	e.CustomFields = v
}

func (eo EventCaseObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(eo.CommonEventCaseObject.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, supportingfunctions.ToStringBeautifulMapSlice(num, eo.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, eo.TagsAll)))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, common.CustomFieldsToStringBeautiful(eo.CustomFields, num)))

	return str.String()
}
