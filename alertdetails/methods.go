package alertdetails

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewEventAlertDetails создаёт новый объект типа EventAlertDetails
func NewEventAlertDetails() *EventAlertDetails {
	return &EventAlertDetails{
		Tags:    make(map[string][]string),
		TagsAll: []string(nil),
	}
}

func (e *EventAlertDetails) Get() *EventAlertDetails {
	return e
}

func (e *EventAlertDetails) GetSourceRef() string {
	return e.SourceRef
}

// SetValueSourceRef значение поля SourceRef
func (e *EventAlertDetails) SetValueSourceRef(v string) {
	e.SourceRef = v
}

// SetAnySourceRef значение поля SourceRef
func (e *EventAlertDetails) SetAnySourceRef(i any) {
	e.SourceRef = fmt.Sprint(i)
}

func (e *EventAlertDetails) GetTitle() string {
	return e.Title
}

// SetValueTitle значение поля Title
func (e *EventAlertDetails) SetValueTitle(v string) {
	e.Title = v
}

// SetAnyTitle значение поля Title
func (e *EventAlertDetails) SetAnyTitle(i any) {
	e.Title = fmt.Sprint(i)
}

func (e *EventAlertDetails) GetDescription() string {
	return e.Description
}

// SetValueDescription значение поля Description
func (e *EventAlertDetails) SetValueDescription(v string) {
	v = strings.ReplaceAll(v, "\t", "")
	v = strings.ReplaceAll(v, "\n", "")

	e.Description = v
}

// SetAny значение поля Description
func (e *EventAlertDetails) SetAnyDescription(i any) {
	str := fmt.Sprint(i)
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	e.Description = str
}

func (e *EventAlertDetails) GetTags() map[string][]string {
	return e.Tags
}

// SetValueTags добаляет значение в Tags по ключу
func (e *EventAlertDetails) SetValueTags(k, v string) bool {
	if _, ok := e.Tags[k]; !ok {
		e.Tags[k] = []string(nil)
	}

	for _, value := range e.Tags[k] {
		if v == value {
			return false
		}
	}

	e.Tags[k] = append(e.Tags[k], v)

	return true
}

// SetAnyTags значение поля Tags
func (e *EventAlertDetails) SetAnyTags(k string, i interface{}) bool {
	return e.SetValueTags(k, fmt.Sprint(i))
}

// SetSliceValueTags значение поля Tags
func (e *EventAlertDetails) SetSliceValueTags(v map[string][]string) {
	e.Tags = v
}

func (a *EventAlertDetails) GetTagsAll() []string {
	return a.TagsAll
}

// SetValueTagsAll значение в список поля TagsAll
func (a *EventAlertDetails) SetValueTagsAll(v string) {
	a.TagsAll = append(a.TagsAll, v)
}

// SetAnyTagsAll значение в список поля TagsAll
func (a *EventAlertDetails) SetAnyTagsAll(i any) {
	a.TagsAll = append(a.TagsAll, fmt.Sprint(i))
}

// ToStringBeautiful форматированный вывод
func (e *EventAlertDetails) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'sourceRef': '%s'\n", ws, e.SourceRef))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws, e.Title))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, e.Description))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, supportingfunctions.ToStringBeautifulMapSlice(num, e.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, e.TagsAll)))

	return str.String()
}
