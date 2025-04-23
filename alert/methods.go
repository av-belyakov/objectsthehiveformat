package alert

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/objectsthehiveformat/alertartifacts"
	"github.com/av-belyakov/objectsthehiveformat/common"
	"github.com/av-belyakov/objectsthehiveformat/supportingfunctions"
)

// NewTypeAlert создаёт новый объект типа TypeAlert
func NewTypeAlert() *TypeAlert {
	return &TypeAlert{
		CommonEventAlertObject: common.CommonEventAlertObject{
			Date:      "1970-01-01T00:00:00+00:00",
			CreatedAt: "1970-01-01T00:00:00+00:00",
			UpdatedAt: "1970-01-01T00:00:00+00:00",
		},
		Tags:         make(map[string][]string),
		TagsAll:      []string(nil),
		CustomFields: common.CustomFields{},
		Artifacts:    make(map[string][]alertartifacts.Artifacts),
	}
}

// Get возвращает объект типа TypeAlert
func (a *TypeAlert) Get() *TypeAlert {
	return a
}

func (a *TypeAlert) GetFollow() bool {
	return a.Follow
}

// SetValueFollow устанавливает значение поля Follow
func (a *TypeAlert) SetValueFollow(v bool) {
	a.Follow = v
}

// SetAnyFollow устанавливает значение поля Follow
func (a *TypeAlert) SetAnyFollow(i any) {
	if v, ok := i.(bool); ok {
		a.Follow = v
	}
}

func (a *TypeAlert) GetSeverity() uint64 {
	return a.Severity
}

// SetValueSeverity устанавливает значение поля Severity
func (a *TypeAlert) SetValueSeverity(v uint64) {
	a.Severity = v
}

// SetAnySeverity устанавливает значение поля Severity
func (a *TypeAlert) SetAnySeverity(i any) {
	if v, ok := i.(float64); ok {
		a.Severity = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		a.Severity = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		a.Severity = v
	}
}

func (a *TypeAlert) GetTags() map[string][]string {
	return a.Tags
}

// SetValueTags добаляет значение в Tags по ключу
func (a *TypeAlert) SetValueTags(k, v string) bool {
	if _, ok := a.Tags[k]; !ok {
		a.Tags[k] = []string(nil)
	}

	v = supportingfunctions.TrimIsNotLetter(v)
	for _, value := range a.Tags[k] {
		if v == value {
			return false
		}
	}

	a.Tags[k] = append(a.Tags[k], v)

	return true
}

// SetAnyTags добаляет значение в Tags по ключу
func (a *TypeAlert) SetAnyTags(k string, i any) bool {
	return a.SetValueTags(k, fmt.Sprint(i))
}

func (a *TypeAlert) GetTagsAll() []string {
	return a.TagsAll
}

// SetValueTagsAll добавляет значение в список поля TagsAll
func (a *TypeAlert) SetValueTagsAll(v string) {
	a.TagsAll = append(a.TagsAll, v)
}

// SetAnyTagsAll добавляет значение в список поля TagsAll
func (a *TypeAlert) SetAnyTagsAll(i any) {
	a.TagsAll = append(a.TagsAll, fmt.Sprint(i))
}

func (a *TypeAlert) GetCustomFields() common.CustomFields {
	return a.CustomFields
}

// SetValueCustomFields устанавливает значение поля CustomFields
func (a *TypeAlert) SetValueCustomFields(v common.CustomFields) {
	a.CustomFields = v
}

func (a *TypeAlert) GetArtifacts() map[string][]alertartifacts.Artifacts {
	return a.Artifacts
}

func (a *TypeAlert) GetKeyArtifacts(k string) ([]alertartifacts.Artifacts, bool) {
	if value, ok := a.Artifacts[k]; ok {
		return value, true
	}

	return nil, false
}

// SetKeyArtifacts добавляет значение в виде среза для поля Artifacts по ключу
func (a *TypeAlert) SetKeyArtifacts(k string, artifacts []alertartifacts.Artifacts) {
	a.Artifacts[k] = artifacts
}

// SetArtifacts устанавливает значение поля Artifacts
func (a *TypeAlert) SetValueArtifacts(v map[string][]alertartifacts.Artifacts) {
	a.Artifacts = v
}

// AddValueArtifact добавляет значение поля Artifacts по ключу
func (a *TypeAlert) AddValueArtifact(k string, v alertartifacts.Artifacts) {
	if _, ok := a.Artifacts[k]; !ok {
		a.Artifacts[k] = []alertartifacts.Artifacts(nil)
	}

	a.Artifacts[k] = append(a.Artifacts[k], v)
}

// ToStringBeautiful форматированный вывод
func (a *TypeAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(a.CommonEventAlertObject.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'follow': '%t'\n", ws, a.Follow))
	str.WriteString(fmt.Sprintf("%s'severity': '%d'\n", ws, a.Severity))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, supportingfunctions.ToStringBeautifulMapSlice(num, a.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, a.TagsAll)))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, common.CustomFieldsToStringBeautiful(a.CustomFields, num)))
	str.WriteString(fmt.Sprintf("%s'artifact': \n%s", ws, func(l map[string][]alertartifacts.Artifacts) string {
		str := strings.Builder{}

		for key, value := range l {
			str.WriteString(fmt.Sprintf("%s%s:\n", supportingfunctions.GetWhitespace(num+1), key))
			for k, v := range value {
				str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+2), k))
				str.WriteString(v.ToStringBeautiful(num + 3))
			}
		}
		return str.String()
	}(a.Artifacts)))

	return str.String()
}
